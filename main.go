package main

import (
	"fmt"
	"flag"
	. "spider/interval/dao"
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
		CreateSlaveServer()
	} else {
		CreateMasterServer()
	}
}