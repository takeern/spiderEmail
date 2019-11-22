package master

import (
	pb "spider/interval/serve/grpc"
	"google.golang.org/grpc"
	"github.com/google/martian/log"
	"spider/interval/conf"
)

type Any interface {}

func CreateConn(ip string) (pb.TaskClient, error) {
	
	conn, err := grpc.Dial(ip + ":" + conf.SLAVE_PORT, grpc.WithInsecure())
	if err != nil {
		log.Errorf("connet to slave node failed, node ip: %s, err: %v", ip, err)
		return nil ,err
	}
	c := pb.NewTaskClient(conn)
	return c, err
}

type Dispatch interface {
	HandleNewIpRegistry(ip string, c pb.TaskClient) (code int, msg string)
	/*
	* 同步数据
	* true - 同步所有改分发器数据
	* false - 仅仅同步变化的数据
	*/
	HandleSyncData(status bool)
}