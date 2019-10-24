package slave

import (
	"log"
	"context"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"spider/interval/conf"
	"net"
	"google.golang.org/grpc"
	pb "spider/interval/serve/grpc"
)

type Server struct{}

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func CreateSlaveServe() {
	lis, err := net.Listen("tcp", ":6011")
	log.Printf("listen: 6011")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTaskServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
	}
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
		err := SendMail(req.EmailInfo.Ac, 
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
		err, emails, urls := SpiderEmail(req.SpiderUrl, 0)
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
	resp, err := http.Get(conf.MASTER_HOST + "/register")
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