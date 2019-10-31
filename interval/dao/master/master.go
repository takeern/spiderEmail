package master

import (
	"github.com/gin-gonic/gin"
	"runtime"
	"spider/interval/conf"
	"spider/interval/dao/utils"
	"strings"
	"strconv"
	"fmt"
)

type MasterServer struct {
	EmailDispatch 		*EmailDispatch
	SpiderDispatch		*SpiderDispatch
	IpList        		map[string]bool
}

type emailInfo struct {
	WaitSpiderLen		int
	SpiderIndex			int
	ErrSpiderLen		int
	SuccessSpider		[]string
	ErrorSpider			[]string
	IpList				[]string
}

type spiderInfo struct {
	WaitSpiderLen		int
	HadSpiderLen		int
	ErroSpiderLen		int
	HostUrl				string
	SpiderTimes			int
	ErroSpider			[]string
	WaitSpider			[]string
	IpList				[]string
}

func (ms *MasterServer)handleIpRegistry(c *gin.Context) {
	ip := c.ClientIP()
	ms.IpList[ip] = true
	taskcode := c.Query("accessTask")
	fmt.Printf(taskcode)
	arr := strings.Split(taskcode, "|")
	fmt.Println(arr)
	if len(arr) == 0 {
		ms.SpiderDispatch.HandleNewIpRegistry(ip)
		ms.EmailDispatch.HandleNewIpRegistry(ip)
		utils.Log.Info("create all task")
	} else {
		for _, item := range arr {
			m, _ := strconv.Atoi(item)
			switch m {
			case conf.SEND_EMAIL:
				utils.Log.Info("create send email task")
				ms.EmailDispatch.HandleNewIpRegistry(ip)
			break;
			case conf.SPIDER_EMAIL:
				utils.Log.Info("create spider email task")
				ms.SpiderDispatch.HandleNewIpRegistry(ip)
			break;
			default:
				utils.Log.Info("unhandle taskcode", item)
			}
		}
	}

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
		"emailInfo": &emailInfo{
			IpList: ms.EmailDispatch.Ip_list,
			WaitSpiderLen: len(ms.EmailDispatch.Email_list),
			SpiderIndex: ms.EmailDispatch.Email_send_index,
			ErrSpiderLen: len(ms.EmailDispatch.Error_Email_list),
			SuccessSpider: ms.EmailDispatch.Success_Email_list,
			ErrorSpider: ms.EmailDispatch.Error_Email_list,
		},
		"spiderInfo": &spiderInfo{
			WaitSpiderLen: ms.SpiderDispatch.Wait_spider_queue.Len(),
			HadSpiderLen: ms.SpiderDispatch.Had_spider_queue.Len(),
			ErroSpiderLen: ms.SpiderDispatch.Error_spider_queue.Len(),
			ErroSpider: ms.SpiderDispatch.Error_spider_queue.Q,
			WaitSpider: ms.SpiderDispatch.Wait_spider_queue.Q,
			HostUrl: ms.SpiderDispatch.Host_url,
			SpiderTimes: ms.SpiderDispatch.Spider_times,
			IpList: ms.SpiderDispatch.Ip_list,
		},
	})
}

func (ms *MasterServer) StarServer() {
	r := gin.Default()
	r.GET("/register", ms.handleIpRegistry)
	r.GET("/getServeInfo", ms.getServeInfo)
	r.Run(":" + conf.HOST_PORT)
}