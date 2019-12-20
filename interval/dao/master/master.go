package master

import (
	"sync"
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
	// switch req.TaskCode {
	// case conf.SYNC_DATA:
	// 	if req.SyncData.SyncType == conf.SYNC_ALL {
	// 		// 同步所有数据, 清理所有数据
	// 		utils.Log.Info("sync all data")
	// 		ms.status.syncIdHistory = ms.status.syncIdHistory[0:0]
	// 		ms.status.syncIdHistory = append(ms.status.syncIdHistory, req.SyncData.SyncId)
	// 		ms.SpiderDispatch.InjectInitData(req.SyncData.SpiderSyncData.SpiderAllData)
	// 	} else if req.SyncData.SyncType == conf.SYNC_RECORD {
	// 		// 同步 record 数据
	// 		i := len(ms.status.syncIdHistory) - 1
	// 		if i >= 0 && ms.status.syncIdHistory[i] == req.SyncData.SyncLastId {	// 保证 顺序一致性
	// 			utils.Log.Info("sync record data")
	// 			ms.status.syncIdHistory = append(ms.status.syncIdHistory, req.SyncData.SyncId)
	// 			ms.SpiderDispatch.InjectRecordData(req.SyncData.SpiderSyncData.SpiderRecordData)
	// 		} else {
	// 			resp.Code = conf.ERROR_SYNCDATA_TASK
	// 		}
	// 	}
	// 	break;
	// default:
	// 	break;
	// }
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
		ms.emailIpList[ip] = true
	} else if s == "spider" {
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
