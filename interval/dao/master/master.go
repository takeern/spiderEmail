package master

import (
	"runtime"
	"strings"
	"strconv"
	"sync"
	// "fmt"

	"spider/interval/conf"
	"spider/interval/modal"
	"spider/interval/net"
	"spider/interval/dao/utils"
	pb "spider/interval/serve/grpc"

	"github.com/gin-gonic/gin"
)

type MasterServer struct {
	mu      			sync.Mutex
	EmailDispatch 		*EmailDispatch
	SpiderDispatch		*SpiderDispatch
	IpList        		map[string]bool
	connClients			map[string]pb.TaskClient
	syncList			*modal.Queue
	status				*IStatus
}

type IStatus struct {
	starSync			bool
	syncId				int64
	sleep				bool
	syncIdHistory		[]int64
}

type emailInfo struct {
	WaitSpiderLen		int
	SpiderIndex			int
	ErrSpiderLen		int
	SuccessSpider		[]string
	ErrorSpider			[]string
	IpList				[]string
}

type spiderInfo struct {
	WaitSpiderLen		int
	HadSpiderLen		int
	ErroSpiderLen		int
	HostUrl				string
	SpiderTimes			int
	ErroSpider			[]string
	HadSpider			[]string
	Wait_spider			[]string
	IpList				[]string
}

func NewMaterServe(sleep bool) *MasterServer {
	ms := &MasterServer{
		IpList:        	make(map[string]bool),
		EmailDispatch: 	CreateEmailDispatch(conf.DB_URL),
		SpiderDispatch: CreateSpiderDispatch(conf.SPIDER_URL, sleep),
		connClients:	make(map[string]pb.TaskClient),
		syncList:		modal.NewQueue(10),
		status:			&IStatus{
			starSync:	false,
			sleep:		sleep,
			syncIdHistory:	make([]int64, 0, 1000),
		},
	}

	// if !sleep {
	// 	ms.StarServer()
	// }

	ms.StarServer()
	return ms
}

// 处理新ip 连接
func (ms *MasterServer)handleIpRegistry(c *gin.Context) {
	var code int
	var msg string
	var grpcC pb.TaskClient
	ip := c.ClientIP()
	taskcode := c.Query("accessTask")

	grpcC, ok := ms.connClients[ip];
	if !ok {
		// 如果无法在缓存中读取 client 
		var err error
		grpcC, err = net.NewClient(ip)
		if err != nil {
			utils.Log.Error("connet to slave node failed, node ip: %s, err: %v", ip, err)
			return
		}
		ms.connClients[ip] = grpcC
	}

	arr := strings.Split(taskcode, "|")
	ms.IpList[ip] = true
	for _, masterIp := range conf.MASTER_IP {
		if masterIp == ip && !ms.syncList.HasValue(ip){
			ms.syncList.Push(ip)
		}
	}

	if len(arr) == 0 {
		code, msg = ms.SpiderDispatch.HandleNewIpRegistry(ip, grpcC)
		codeOther, msgOther := ms.EmailDispatch.HandleNewIpRegistry(ip, grpcC)
		if !(code == conf.RegisterCodeSuccess && codeOther == conf.RegisterCodeSuccess) {
			code = conf.RegisterCodeError
			msg += msgOther
		}
		utils.Log.Info("create all task")
	} else {
		for _, item := range arr {
			m, _ := strconv.Atoi(item)
			switch m {
			case conf.SEND_EMAIL:
				utils.Log.Info("create send email task")
				code, msg = ms.EmailDispatch.HandleNewIpRegistry(ip, grpcC)
			break;
			case conf.SPIDER_EMAIL:
				utils.Log.Info("create spider email task")
				code, msg = ms.SpiderDispatch.HandleNewIpRegistry(ip, grpcC)
			break;
			default:
				utils.Log.Info("unhandle taskcode", item)
			}
		}
	}

	// 如果注册成功并且存在可以同步的节点，同时还未开启同步则开启同步
	if (code == conf.RegisterCodeSuccess && ms.syncList.Len() != 0 && !ms.status.starSync) {
		ms.StarSyncData()
		ms.status.starSync = true
	}

	c.JSON(200, gin.H{
		"code": code,
		"msg": msg,
	})
}

func (ms *MasterServer) StarServer() {
	r := gin.Default()
	r.GET("/register", ms.handleIpRegistry)
	r.GET("/getServeInfo", ms.getServeInfo)
	r.Run(":" + conf.HOST_PORT)
}

func (ms *MasterServer) getServeInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": conf.RegisterCodeSuccess,
		"ip_list": ms.IpList,
		"Goroutines": runtime.NumGoroutine(),
		"emailInfo": &emailInfo{
			IpList: ms.EmailDispatch.Ip_list.Q,
			WaitSpiderLen: len(ms.EmailDispatch.Email_list),
			SpiderIndex: ms.EmailDispatch.Email_send_index,
			ErrSpiderLen: len(ms.EmailDispatch.Error_Email_list),
			SuccessSpider: ms.EmailDispatch.Success_Email_list,
			ErrorSpider: ms.EmailDispatch.Error_Email_list,
		},
		"spiderInfo": &spiderInfo{
			WaitSpiderLen: ms.SpiderDispatch.Data.Wait_spider_queue.Len(),
			HadSpiderLen: ms.SpiderDispatch.Data.Had_spider_queue.Len(),
			ErroSpiderLen: ms.SpiderDispatch.Data.Error_spider_queue.Len(),
			ErroSpider: ms.SpiderDispatch.Data.Error_spider_queue.Q,
			HadSpider: ms.SpiderDispatch.Data.Had_spider_queue.Q,
			Wait_spider: ms.SpiderDispatch.Data.Wait_spider_queue.Q,
			HostUrl: ms.SpiderDispatch.Data.Host_url,
			SpiderTimes: ms.SpiderDispatch.Data.Spider_times,
			IpList: ms.SpiderDispatch.Data.Ip_list.Q,
		},
	})
}

// 睡眠状态下 主节点同步数据
func (ms *MasterServer) HandleReq(req *pb.HandleTaskReq) *pb.HandleTaskResp{
	resp := &pb.HandleTaskResp {
		Code: conf.SUCCESS_TASK,
	}
	switch req.TaskCode {
	case conf.SYNC_DATA:
		list := ms.status.syncIdHistory
		if req.SyncData.SyncType == conf.SYNC_ALL {
			// 同步所有数据, 清理所有数据
			utils.Log.Info("sync all data")
			list = list[0:0]
			list = append(list, req.SyncData.SyncId)
			ms.SpiderDispatch.InjectInitData(req.SyncData.SpiderSyncData.SpiderAllData)
		} else if req.SyncData.SyncType == conf.SYNC_RECORD {
			// 同步 record 数据
			if list[len(list) - 1] == req.SyncData.SyncLastId {	// 保证 顺序一致性
				utils.Log.Info("sync record data")
				list = append(list, req.SyncData.SyncId)
				ms.SpiderDispatch.InjectRecordData(req.SyncData.SpiderSyncData.SpiderRecordData)
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