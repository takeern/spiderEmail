package dao

import (
	"golang.org/grpc-go"
	"log"
	"golang.org/x/net/context"
	"net"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"spider/interval/conf"
	pb "spider/interval/serve/grpc"
	"runtime"
	"spider/interval/dao/master"
	"spider/interval/dao/slave"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	registerIp(0)
}

func (s *Server) HandleTask(ctx context.Context, req *pb.HandleTaskReq) (*pb.HandleTaskResp, error) {
	resp := new(pb.HandleTaskResp)
	selectTaskRun(req, resp)
	return resp, nil
}

/*
 * 选择任务类型
 * 1000 发送 email
 * 1001 爬取页面 email
 */
func selectTaskRun(req *pb.HandleTaskReq, resp *pb.HandleTaskResp) (*pb.HandleTaskResp) {
	switch req.TaskCode {
	case conf.SEND_EMAIL:
		err := slave.SendMail(req.EmailInfo.Ac,
			req.EmailInfo.Ps,
			req.EmailInfo.Host,
			req.EmailInfo.Receive,
			"无主题",
			conf.EmailModalList[req.EmailInfo.ModalIndex],
			"html")
		if err != nil {
			resp.Code = 10001
			resp.ErrorMsg = err.Error()
		} else {
			resp.Code = 10000
		}

		return resp
	case conf.SPIDER_EMAIL:
		err, emails, urls := slave.SpiderEmail(req.SpiderUrl, 0)
		if err != nil {
			resp.Code = 10002
			resp.ErrorMsg = err.Error()
		} else {
			resp.Code = 10000
			resp.SpiderInfo.Emails = emails
			resp.SpiderInfo.Urls = urls
		}
		return resp
	default:
		resp.Code = 10003
		resp.ErrorMsg = "unhandle task code"
		return resp
	}
}


func registerIp(times int) {
	if (times > conf.RETRY_REGISTER_TIMES) {
		log.Fatal("register many time exit")
	}
	resp, err := http.Get(conf.MASTER_HOST + "?token=" + conf.MASTER_TOKEN)
	if err != nil {
		times ++
		registerIp(times)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		times ++
		registerIp(times)
		return
	}
	if (body[0] == 0) {
		log.Println("register success")
	} else {
		log.Fatal("register res unhandle %v", body)
	}
}


////////////////////////////////////////////////


func CreateSlaveServer() {
	lis, err := net.Listen("tcp", ":" + conf.SLAVE_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listen: " + conf.SLAVE_PORT + " port succeed")

	s := grpc.NewServer()
	pb.RegisterTaskServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}




/*---------------------------------------------------------*/
/*---------------------------------------------------------*/
/*---------------------------------------------------------*/
/*---------------------------------------------------------*/
/*---------------------------------------------------------*/


type MasterServer struct {
	emailDispatch	    *master.Dispatch
	ipList		        map[string]bool
}

func (ms *MasterServer)handleIpRegistry(c *gin.Context) {
	ip := c.ClientIP()
	ms.ipList[ip] = true
	ms.emailDispatch.HandleNewIpRegistry(ip)

	c.JSON(200, gin.H{
		"code": "10000",
		"msg": "ip registry success",
	})
}

func (ms *MasterServer) getServeInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "10000",
		"ip_list": ms.ipList,
		"Goroutines": runtime.NumGoroutine(),
		"waitSpiderLen": len(ms.emailDispatch.Email_list),
		"spiderIndex": ms.emailDispatch.Email_send_index,
		"errSpiderLen": len(ms.emailDispatch.Error_Email_list),
		"successSpider": ms.emailDispatch.Success_Email_list,
		"errorSpider": ms.emailDispatch.Error_Email_list,
	})
}

func (ms *MasterServer)StarServer() {
	r := gin.Default()
	r.GET("/register", ms.handleIpRegistry)
	r.GET("/getServeInfo", ms.getServeInfo)
	r.Run(":" + conf.HOST_PORT)
}

func CreateMasterServer() *MasterServer {

	ms := &MasterServer{
		ipList:             make(map[string]bool),
		emailDispatch:		master.CreateEmailDispatch(conf.DB_URL),
	}

	ms.StarServer()

	return ms
}