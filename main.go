package main

import (
	"spider/interval/dao"
)



func main() {
	// dao.GetUrl("http://www.ijetch.org/index.php?m=content&c=index&a=show&catid=103&id=1372")
	// dao.Log.Debug("ddadsasd")
	// mb := dao.NewDb("http://www.ijetch.org/")
	// mb.InsertData("http://www.ijetch.org/", "hasdhas")
	dao.GoSend()
	// dao.GetUrl("http://www.ijetch.org")
	// dao.GetUrl("http://www.ijetch.org/index.php?m=content\u0026c=index\u0026a=show\u0026catid=103\u0026id=1373")
}