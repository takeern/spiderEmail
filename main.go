package main

import (
	"flag"
	"fmt"
	. "spider/interval/dao"
)

var (
	help     bool
	master   bool
	slave    bool
)

func init() {
	flag.BoolVar(&help, "help", false, "show usage")
	flag.BoolVar(&master, "master", false, "start as a master node")
	flag.BoolVar(&slave, "slave", false, "start as a slave node")
}

func showUsage()  {
	fmt.Print("Usage:   ./spider master --- start as a master node\n" +
		"         ./spider slave --- start as a slave node\n\n" +
		"         good luck, my bro\n")
}

func main() {
	flag.Parse()

	if help {
		showUsage()
		return
	}

	if slave {
		CreateSlaveServer()
		return
	}

	if master {
		CreateMasterServer()
		return
	}

	showUsage()
}