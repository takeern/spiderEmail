package master

import (
	"github.com/gin-gonic/gin"
	"runtime"
	"spider/interval/conf"
)

type MasterServer struct {
	EmailDispatch *Dispatch
	IpList        map[string]bool
}

func (ms *MasterServer)handleIpRegistry(c *gin.Context) {
	ip := c.ClientIP()
	ms.IpList[ip] = true
	ms.EmailDispatch.HandleNewIpRegistry(ip)

	c.JSON(200, gin.H{
		"code": "10000",
		"msg": "ip registry success",
	})
}

func (ms *MasterServer) getServeInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "10000",
		"ip_list": ms.IpList,
		"Goroutines": runtime.NumGoroutine(),
		"waitSpiderLen": len(ms.EmailDispatch.Email_list),
		"spiderIndex": ms.EmailDispatch.Email_send_index,
		"errSpiderLen": len(ms.EmailDispatch.Error_Email_list),
		"successSpider": ms.EmailDispatch.Success_Email_list,
		"errorSpider": ms.EmailDispatch.Error_Email_list,
	})
}

func (ms *MasterServer) StarServer() {
	r := gin.Default()
	r.GET("/register", ms.handleIpRegistry)
	r.GET("/getServeInfo", ms.getServeInfo)
	r.Run(":" + conf.HOST_PORT)
}