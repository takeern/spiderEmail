package dao

import (
	// "fmt"
	"runtime"
	"spider/interval/dao/master"
	"spider/interval/conf"
	"spider/interval/dao/slave"
	"github.com/gin-gonic/gin"
)

var (
	S		*Serve
)

type Serve struct {
	Ms 			*master.MasterServe
	Ip_list		map[string]bool
	status		string
}

type EmailInfo struct {
	waitSpiderLen			int
	spiderIndex  			int
	errSpiderLen  			int
	successSpider  			[]string
	errorSpider  			[]string
}

func CreateServe(nodeType string) *Serve {
	var typeStatus bool
	if nodeType == "master" {
		typeStatus = true
	} else {
		slave.CreateSlaveServe()
		typeStatus = false
	}
	S = &Serve{
		Ip_list: make(map[string]bool),
		Ms:		master.NewMasterServe(typeStatus),
		status: nodeType,
	}

	if S.status == "master" {
		StarServe()
	}

	return S
}


func (s *Serve)handleIpRegistry(c *gin.Context) {
	ip := c.ClientIP()
	s.Ip_list[ip] = true
	if s.status == "master" {
		s.Ms.HandleNewIpRegistry(ip)
	}
	c.JSON(200, gin.H{
		"code": "10000",
		"msg": "ip registry success",
	})
}

func (s *Serve) getServeInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "10000",
		"ip_list": s.Ip_list,
		"Goroutines": runtime.NumGoroutine(),
		"waitSpiderLen": len(s.Ms.Email.Email_list),
		"spiderIndex": s.Ms.Email.Email_send_index,
		"errSpiderLen": len(s.Ms.Email.Error_Email_list),
		"successSpider": s.Ms.Email.Success_Email_list,
		"errorSpider": s.Ms.Email.Error_Email_list,
	})
}

func StarServe() {
	r := gin.Default()
	r.GET("/register", S.handleIpRegistry)
	r.GET("/getServeInfo", S.getServeInfo)
	r.Run(":" + conf.HOST_PORT)
}