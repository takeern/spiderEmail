package dao

import (
	"spider/interval/dao/master"
	"spider/interval/dao/slave"
	"spider/interval/net"
	"spider/interval/conf"
)

func CreateMasterServer() {
	master.NewMaterServe(false)
}

func CreateSlaveServer() {
	rpc := net.NewServer()
	ms := master.NewMaterServe(true)
	rpc.AddListener(slave.HandleReq, conf.TYPE_SLAVE)
	rpc.AddListener(ms.HandleReq, conf.TYPE_MASTER)
	rpc.Listen()
}
