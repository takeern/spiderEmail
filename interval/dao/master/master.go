package master

import (
	"sync"
	"encoding/json"
	"context"
	"time"
	"io/ioutil"
	// "log"
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
	urlRecord			[]string
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
			urlRecord: make([]string, 0, 100),
		},
	}

	if !sleep {
		ms.StarServer()
	}
	
	return ms
}

func (ms *MasterServer) getSnapchat() []string{
	files, _ := ioutil.ReadDir("./snapchat")
	names := make([]string, 0, 100)
	for _, f := range files {
		names = append(names, f.Name())
	}
	return names
}

// 物理方式（snapchat）存储 当前服务状态
func (ms *MasterServer) setSnapchat() {
	data := ms.getAllSyncData(false)
	b, _ := json.Marshal(data)
	name := "./snapchat/" + time.Now().Format("2006-01-02-15:04:05") +  ".snapchat"
	err := ioutil.WriteFile(name, b, 0666)
	if err != nil {
		utils.Log.Error("存档失败", name, err)
	} else {
		utils.Log.Info("存档成功", name)
	}
}

func (ms *MasterServer) useSnapchat(name string) (int32, string) {
	data, err := ioutil.ReadFile("./snapchat/" + name)
	if err != nil {
		return conf.ERROR_REP, conf.SNAPCHAT_ERROR + " can not find"
	}
	var req *pb.HandleTaskReq
	if err := json.Unmarshal(data, &req); err != nil {
		return conf.ERROR_REP, conf.SNAPCHAT_ERROR + " unmarshal error"
	}
	ms.HandleReq(req)
	return conf.SUCCESS_TASK, conf.SNAPCHAT_SUCCESS
}

// 创建一个 爬虫分发器
func (ms *MasterServer) creatSpider(url string) (int, string) {
	code := conf.ERROR_SPIDER_TASK
	msg := conf.CreateSpiderURLError;
	utils.Log.Info("处理 创建新的 爬虫任务: %s", url)
	if url != "" {
		_, ok := ms.SpiderDispatchs[url]
		if !ok {
			code = conf.SUCCESS_TASK
			d := d.NewSpider(ms.ctx, url, false, ms.spiderIpList, ms.connClients)
			ms.SpiderDispatchs[url] = d
			msg = conf.CreateSpiderMsgSuccess
		} else {
			msg = conf.CreateSpiderURLRepeat
		}
	}
	return code, msg
}

// 删除某项 spider
func (ms *MasterServer) DeleteSpider(url string) {
	utils.Log.Info("删除某项 spider: %s", url)
	delete(ms.SpiderDispatchs, url)
}

// 获取 ip 列表
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
func (ms *MasterServer) getSyncData(record bool) (map[string]string, []string){
	r := make(map[string]string)
	urls := make([]string, 1, 30)
	if record {
		for k, v := range ms.status.ipRecord {
			r[k] = v
			delete(ms.status.ipRecord, k)
		}
		copy(urls, ms.status.urlRecord)
		ms.status.urlRecord = ms.status.urlRecord[0:0]
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
		for url, _ := range ms.SpiderDispatchs {
			urls = append(urls, url)
		}
	}
	return r, urls
}