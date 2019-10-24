package main

import (
	// "net"
	// "log"
	// "google.golang.org/grpc"

	// pb "spider/interval/serve/grpc"
	"spider/interval/dao"
)


func main() {
	dao.CreateServe("slave")
}


// func main() {
// 	lis, err := net.Listen("tcp", ":6011")
// 	log.Printf("listen: 6011")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterTaskServer(s, &dao.Server{})
// 	if err := s.Serve(lis); err != nil {
//         log.Fatalf("failed to serve: %v", err)
// 	}
// }
