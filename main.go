package main

import (
	"fmt"
	"flag"
	"log"
	"net"
	"spider/interval/dao"
	"golang.org/grpc-go"
	pb "spider/interval/serve/grpc"
)


func main() {
	help := flag.Bool("help", false, "show usage")
	slave := flag.Bool("slave", false, "start as a slave node")
	flag.Parse()

	if *help {
		fmt.Print("Usage:   ./spider --- start as a master node\n" +
			"         ./spider slave --- start as a slave node\n")
		return
	}

	if *slave {
		dao.CreateSlaveServer()
	} else {
		dao.CreateMasterServer()
	}
}