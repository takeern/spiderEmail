package dao

import (
	// "golang.org/grpc-go"
	"google.golang.org/grpc"
	"log"
	"net"
	"spider/interval/conf"
	pb "spider/interval/serve/grpc"
	"spider/interval/dao/master"
	"spider/interval/dao/slave"
)

func CreateMasterServer() {

	ms := &master.MasterServer{
		IpList:        make(map[string]bool),
		EmailDispatch: master.CreateEmailDispatch(conf.DB_URL),
	}

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