package dao

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"spider/interval/conf"
	pb "spider/interval/serve/grpc"
	"spider/interval/dao/master"
	"spider/interval/dao/slave"
)

func CreateMasterServer() {
	ms := master.NewMaterServe()

	ms.StarServer()
}

func CreateSlaveServer() {
	lis, err := net.Listen("tcp", ":" + conf.SLAVE_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listen: " + conf.SLAVE_PORT + " port succeed")

	s := grpc.NewServer()
	pb.RegisterTaskServer(s, &slave.SlaveServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	slave.RegisterIp(0)
}
