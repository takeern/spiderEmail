package master

import (
	pb "spider/interval/serve/grpc"
	"google.golang.org/grpc"
	"github.com/google/martian/log"
	"spider/interval/conf"
)

func CreateConn(ip string) (pb.TaskClient, error) {
	
	conn, err := grpc.Dial(ip + ":" + conf.SLAVE_PORT, grpc.WithInsecure())
	if err != nil {
		log.Errorf("connet to slave node failed, node ip: %s, err: %v", ip, err)
		return nil ,err
	}
	c := pb.NewTaskClient(conn)
	return c, err
}