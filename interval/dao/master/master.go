package master

import (
	"sync"
	"context"
	"log"
	// "fmt"

	"spider/interval/conf"
	"spider/interval/modal"
	"spider/interval/dao/utils"
	d "spider/interval/dao/dispatch"
	pb "spider/interval/serve/grpc"

)

type MasterServer struct {
	mu      			sync.Mutex
	EmailDispatch 		*EmailDispatch
	SpiderDispatchs		map[string]*d.Spider
	spiderIpList        map[string]bool
	emailIpList         map[string]bool
	connClients			map[string]pb.TaskClient
	syncList			*modal.Queue
	status				*IStatus
	ctx					context.Context
}

type IStatus struct {
	starSync			bool
	syncId				int64
	sleep				bool
	syncIdHistory		[]int64
	ipRecord			map[string]string
}

func NewMaterServe(sleep bool) *MasterServer {
	ms := &MasterServer{
		spiderIpList:   make(map[string]bool),
		emailIpList:    make(map[string]bool),
		EmailDispatch: 	CreateEmailDispatch(conf.DB_URL),
		connClients:	make(map[string]pb.TaskClient),
		SpiderDispatchs:make(map[string]*d.Spider),
		syncList:		modal.NewQueue(10),
		ctx:			context.Background(),
		status:			&IStatus{
			starSync:	false,
			sleep:		sleep,
			syncIdHistory:	make([]int64, 0, 1000),
			ipRecord:	make(map[string]string),
		},
	}

	if !sleep {
		ms.StarServer()
	}
	
	return ms
}

// 睡眠状态下 主节点同步数据
func (ms *MasterServer) HandleReq(req *pb.HandleTaskReq) *pb.HandleTaskResp{
	resp := &pb.HandleTaskResp {
		Code: conf.SUCCESS_TASK,
	}
	switch req.TaskCode {
	case conf.SYNC_DATA:
		if req.SyncData.SyncType == conf.SYNC_ALL {
			// 同步所有数据, 清理所有数据
			ms.status.syncIdHistory = ms.status.syncIdHistory[0:0]
			ms.status.syncIdHistory = append(ms.status.syncIdHistory, req.SyncData.SyncId)
			// ms.SpiderDispatch.InjectInitData(req.SyncData.SpiderSyncData.SpiderAllData) //to do add inject data
		} else if req.SyncData.SyncType == conf.SYNC_RECORD {
			// 同步 record 数据
			i := len(ms.status.syncIdHistory) - 1
			if i >= 0 && ms.status.syncIdHistory[i] == req.SyncData.SyncLastId {	// 保证 顺序一致性
				utils.Log.Info("sync record data")
				ms.status.syncIdHistory = append(ms.status.syncIdHistory, req.SyncData.SyncId)
				// ms.SpiderDispatch.InjectRecordData(req.SyncData.SpiderSyncData.SpiderRecordData) //to do add inject data
			} else {
				resp.Code = conf.ERROR_SYNCDATA_TASK
			}
		}
		break;
	default:
		break;
	}
	return resp
}

// 删除某项 spider
func (ms *MasterServer) DeleteSpider(url string) {
	utils.Log.Info("删除某项 spider: %s", url)
	delete(ms.SpiderDispatchs, url)
}

// 删除某项 spider
func (ms *MasterServer) GetIplist(s string) map[string]bool {
	if s == "email" {
		return ms.emailIpList
	}
	return ms.spiderIpList
}

func (ms *MasterServer) SetIplist(s string, ip string) {
	utils.Log.Info("添加 ip list 类型: %d, ip: ", s, ip)
	if s == "email" {
		if v, ok := ms.status.ipRecord[ip]; ok && v == conf.IP_SPIDER {
			ms.status.ipRecord[ip] = conf.IP_ALL
		} else {
			ms.status.ipRecord[ip] = conf.IP_EMAIL
		}
		ms.emailIpList[ip] = true
	} else if s == "spider" {
		if v, ok := ms.status.ipRecord[ip]; ok && v == conf.IP_EMAIL {
			ms.status.ipRecord[ip] = conf.IP_ALL
		} else {
			ms.status.ipRecord[ip] = conf.IP_SPIDER
		}
		ms.spiderIpList[ip] = true
	}
}

// 删除某项 rpc client
func (ms *MasterServer) DeleteConnClient(ip string) {
	utils.Log.Info("删除某项 rpc client: %s", ip)
	delete(ms.connClients, ip)
}

// 获取 rpc client
func (ms *MasterServer) GetConnClients() map[string]pb.TaskClient {
	return ms.connClients
}

// set rpc client
func (ms *MasterServer) SetConnClients(ip string, c pb.TaskClient) {
	utils.Log.Info("创建 新的 grpc 句柄, ip: ", ip)
	ms.connClients[ip] = c
}

/*
 * 同步数据
 * true 状态下仅同步本次时间段内 变化
 * false 情况下 获取所有需要的数据
 */
func (ms *MasterServer) getSyncData(record bool) (map[string]string){
	r := make(map[string]string)
	if record {
		for k, v := range ms.status.ipRecord {
			r[k] = v
			delete(ms.status.ipRecord, k)
		}
	} else {
		for ip, _ := range ms.spiderIpList {
			r[ip] = conf.IP_SPIDER
		}

		for ip, _ := range ms.emailIpList {
			if _, ok := r[ip]; ok {
				r[ip] = conf.IP_ALL
			} else {
				r[ip] = conf.IP_EMAIL
			}
		}
	}
	return r
}