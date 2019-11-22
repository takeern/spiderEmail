package master

import (
	"github.com/gin-gonic/gin"
	"runtime"
	"spider/interval/conf"
	"spider/interval/modal"
	"spider/interval/dao/utils"
	"strings"
	"strconv"
	"time"
	"context"
	"sync"
	pb "spider/interval/serve/grpc"
	"google.golang.org/grpc"
)

type MasterServer struct {
	mu      			sync.Mutex
	EmailDispatch 		*EmailDispatch
	SpiderDispatch		*SpiderDispatch
	IpList        		map[string]bool
	connClients			map[string]pb.TaskClient
	syncId				int
	syncList			*modal.Queue
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

func NewMaterServe() *MasterServer {
	ms := &MasterServer{
		IpList:        	make(map[string]bool),
		EmailDispatch: 	CreateEmailDispatch(conf.DB_URL),
		SpiderDispatch: CreateDispatchSpider(conf.SPIDER_URL),
		connClients:	make(map[string]pb.TaskClient),
		syncList:		modal.NewQueue(10),
	}
	return ms
}

func (ms *MasterServer) CreateConn(ip string) (pb.TaskClient, error) {
	// 如果存在直接在缓存中获取
	if c, ok := ms.connClients[ip]; ok {
		return c, nil
	}

	conn, err := grpc.Dial(ip + ":" + conf.SLAVE_PORT, grpc.WithInsecure())
	c := pb.NewTaskClient(conn)
	ms.connClients[ip] = c

	return c, err
}

func (ms *MasterServer)handleIpRegistry(c *gin.Context) {
	var code int
	var msg string
	ip := c.ClientIP()
	taskcode := c.Query("accessTask")

	grpcC, err := ms.CreateConn(ip)
	if err != nil {
		utils.Log.Error("connet to slave node failed, node ip: %s, err: %v", ip, err)
		return
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

func (ms *MasterServer) SyncData() {
	go func ()  {
		for {
			ms.mu.Lock()
			// 将 sync ip 同步数据
			if (ms.syncList.Len() != 0) {
				firstSend := true
				if ms.syncId != 0 {
					firstSend = false
				}
				req := ms.getSyncData(firstSend)

				for _, masterIp := range ms.syncList.Q {
					c, ok := ms.connClients[masterIp]
					if ok {
						resp, err := c.HandleTask(context.Background(), req)
						if err != nil {
							utils.Log.Error("grpc: sync data error ", err)
						} else if resp.Code != 10000 {
							utils.Log.Info("grpc: sync data get init data", err)
							initReq := ms.getSyncData(true)
							c.HandleTask(context.Background(), initReq)
						}
					}
				}
			}
			ms.mu.Unlock()
			time.Sleep(conf.WAIT_SYNC_DATA * time.Second)
		}
	}()
}

func (ms *MasterServer) getSyncData(status bool) *pb.HandleTaskReq {
	var spiderData SpiderData
	spiderData = ms.SpiderDispatch.HandleSyncData(status)

	req := &pb.HandleTaskReq {
		TaskCode: conf.SYNC_DATA,
		SyncData: &pb.SyncData{
			SyncSpiderData: &pb.SyncSpiderData{
				IpList: spiderData.Ip_list.Q,
				CloseIpList: spiderData.Close_ip_list.Q,
				WaitSpiderQueue: spiderData.Wait_spider_queue.Q,
				HadSpiderQueue: spiderData.Had_spider_queue.Q,
				ErrorSpiderQueue: spiderData.Error_spider_queue.Q,
				CacheEmail: spiderData.Cache_email,
			},
		},
	}

	return req
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