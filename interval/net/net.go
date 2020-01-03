package net

import (
	"context"
	"crypto/tls"
	"net/http"
	"net"
	"log"
	"spider/interval/conf"

	pb "spider/interval/serve/grpc"
	"google.golang.org/grpc"
)

type RpcServe struct {
	masterHandle	handleReq
	slaveHandle		handleReq
}

type handleReq func(req *pb.HandleTaskReq) *pb.HandleTaskResp

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

// 创建 rpc 服务端
func NewServer() (*RpcServe) {
	rpc := &RpcServe{}
	return rpc
}

func (rpc *RpcServe) Listen() {
	lis, err := net.Listen("tcp", ":" + conf.SLAVE_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listen: " + conf.SLAVE_PORT + " port succeed")

	s := grpc.NewServer()
	pb.RegisterTaskServer(s, rpc)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// 创建 rpc 客户端
func NewClient(ip string) (pb.TaskClient, error) {
	conn, err := grpc.Dial(ip + ":" + conf.SLAVE_PORT, grpc.WithInsecure())
	c := pb.NewTaskClient(conn)

	return c, err
}

// 添加 request 请求监听
func (rpc *RpcServe) AddListener(fn handleReq, s string) {
	if s == conf.TYPE_MASTER {
		rpc.masterHandle = fn
	} else if s == conf.TYPE_SLAVE {
		rpc.slaveHandle = fn
	}
}


func (rpc *RpcServe) HandleTask(ctx context.Context, req *pb.HandleTaskReq) (*pb.HandleTaskResp, error) {
	resp := &pb.HandleTaskResp{}

	if	req.TaskCode > conf.TASK_BOUNDARY {
		// handle master task
		if rpc.masterHandle != nil {
			resp = rpc.masterHandle(req)
		} else {
			resp.Code = conf.ERROR_MASTER_TASK
			resp.ErrorMsg = "no master req listener"
		}
	} else {
		// handle slave task
		if rpc.slaveHandle != nil {
			resp = rpc.slaveHandle(req)
		} else {
			resp.Code = conf.ERROR_SLAVE_TASK
			resp.ErrorMsg = "no slave req listener"
		}
	}
	return resp, nil
}