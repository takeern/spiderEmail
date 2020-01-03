package master

import (
	"runtime"
	"strings"
	"strconv"
	// "fmt"

	"spider/interval/conf"
	"spider/interval/net"
	"spider/interval/dao/utils"
	pb "spider/interval/serve/grpc"

	"github.com/gin-gonic/gin"
)

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
	HadSpider			[]string
	Wait_spider			[]string
	IpList				[]string
}

// 处理新ip 连接
func (ms *MasterServer)handleIpRegistry(c *gin.Context) {
	ip := c.ClientIP()
	taskcode := c.Query("accessTask")
	code, msg := ms.register(ip, taskcode)

	c.JSON(200, gin.H{
		"code": code,
		"msg": msg,
	})
}

func (ms *MasterServer) StarServer() {
	r := gin.Default()
	r.GET("/register", ms.handleIpRegistry)
	r.GET("/getServeInfo", ms.getServeInfo)
	r.GET("/setSlaveIp", ms.hanleSetSlaveIp)
	r.GET("/createSpider", ms.handleCreateSpider)
	r.GET("/deleteSpider", ms.handleDeleteSpider)
	r.GET("/setSnapchat", ms.handleSetSnapchat)
	r.GET("/getSnapchat", ms.handleGetSnapchat)
	r.GET("/useSnapchat", ms.handleUseSnapchat)
	if ms.status.sleep {
		r.Run(":" + conf.SLAVE_PORT)
	} else {
		r.Run(":" + conf.HOST_PORT)
	}
}

func (ms *MasterServer) getServeInfo(c *gin.Context) {
	url := c.Query("url")
	info := &spiderInfo{}
	d, ok := ms.SpiderDispatchs[url]
	if ok {
		info = &spiderInfo{
			WaitSpiderLen: d.Data.Wait_spider_queue.Len(),
			HadSpiderLen: d.Data.Had_spider_queue.Len(),
			ErroSpiderLen: d.Data.Error_spider_queue.Len(),
			ErroSpider: d.Data.Error_spider_queue.Q,
			HadSpider: d.Data.Had_spider_queue.Q,
			Wait_spider: d.Data.Wait_spider_queue.Q,
			HostUrl: d.Data.Host_url,
			SpiderTimes: d.Data.Spider_times,
		}
	}
	urls := make([]string, 0, 100)
	for s, _ := range ms.SpiderDispatchs {
		urls = append(urls, s)
	}
	c.JSON(200, gin.H{
		"code": conf.RegisterCodeSuccess,
		"emailIpList": ms.emailIpList,
		"spiderIpList": ms.spiderIpList,
		"Goroutines": runtime.NumGoroutine(),
		"spdierUrl": urls,
		"spiderInfo": info,
	})
}

/*
 * 处理 创建新的 爬虫任务
 * 检查当前爬虫是否含有该 目标
 */
 func (ms *MasterServer) handleCreateSpider(c *gin.Context) {
	url := c.Query("url")
	code, msg := ms.creatSpider(url)

	c.JSON(200, gin.H{
		"code": code,
		"msg": msg,
	})
}

/*
 * 处理 删除 爬虫任务
 * 检查当前爬虫是否含有该 目标
 */
func (ms *MasterServer) handleDeleteSpider(c *gin.Context) {
	url := c.Query("url")
	code := conf.ERROR_SPIDER_TASK
	msg := conf.DeleteSpiderURLError;
	utils.Log.Info("处理 删除 爬虫任务: %s", url)
	if url != "" {
		d, ok := ms.SpiderDispatchs[url]
		if ok {
			d.Exit()
			ms.DeleteSpider(url)
			code = conf.SUCCESS_TASK
			msg = conf.DeleteSpiderMsgSuccess
		}
	}

	c.JSON(200, gin.H{
		"code": code,
		"msg": msg,
	})
}

func (ms *MasterServer) hanleSetSlaveIp(c *gin.Context) {
	ip := c.Query("ip")
	taskcode := c.Query("accessTask")
	var code int
	var msg string
	if len(ip) == 0 || len(taskcode) == 0 {
		code = conf.RegisterCodeError
		msg = conf.RegisterMsgErrorNoInput
	} else {
		code, msg = ms.register(ip, taskcode)
	}

	c.JSON(200, gin.H{
		"code": code,
		"msg": msg,
	})
}

func (ms *MasterServer)register(ip string, taskcode string) (code int, msg string){
	var grpcC pb.TaskClient
	connc := ms.GetConnClients()
	grpcC, ok := connc[ip]
	if !ok {
		// 如果无法在缓存中读取 client 
		var err error
		grpcC, err = net.NewClient(ip)
		if err != nil {
			utils.Log.Error("connet to slave node failed, node ip: %s, err: %v", ip, err)
			return
		}
		ms.SetConnClients(ip, grpcC)
	}

	arr := strings.Split(taskcode, "|")
	for _, masterIp := range conf.MASTER_IP {	// 判断该ip 是否开启同步
		if masterIp == ip && !ms.syncList.HasValue(ip){
			ms.syncList.Push(ip)
		}
	}

	for _, item := range arr {
		m, _ := strconv.Atoi(item)
		switch m {
		case conf.SEND_EMAIL:
			if _, ok := ms.GetIplist("email")[ip]; ok {
				msg += conf.RegisterMsgErrorRepeat
				code = conf.RegisterCodeError
			} else {
				ms.SetIplist("email", ip)
				msg += conf.RegisterMsgSuccess
				code = conf.RegisterCodeSuccess
			}
			msg += "email"
			break;
		case conf.SPIDER_EMAIL:
			if _, ok := ms.GetIplist("spider")[ip]; ok {
				msg += conf.RegisterMsgErrorRepeat
				code = conf.RegisterCodeError
			} else {
				ms.SetIplist("spider", ip)
				msg += conf.RegisterMsgSuccess
				code = conf.RegisterCodeSuccess
			}
			msg += "spider"
			break;
		default:
			utils.Log.Info("unhandle taskcode", item)
		}
	}

	// 如果注册成功并且存在可以同步的节点，同时还未开启同步则开启同步
	if (code == conf.RegisterCodeSuccess && ms.syncList.Len() != 0 && !ms.status.starSync) {
		ms.StarSyncData()
		ms.status.starSync = true
	}
	return code, msg
}

// 物理方式（snapchat）存储 当前服务状态
func (ms *MasterServer) handleSetSnapchat(c *gin.Context) {
	ms.setSnapchat()
}

// 获取当前所有 snapchat
func (ms *MasterServer) handleGetSnapchat(c *gin.Context) {
	names := ms.getSnapchat()
	c.JSON(200, gin.H{
		"code": conf.SUCCESS_TASK,
		"data": names,
	})
}

// 使用 snapchat
func (ms *MasterServer) handleUseSnapchat(c *gin.Context) {
	name := c.Query("name")
	code, msg := ms.useSnapchat(name)
	c.JSON(200, gin.H{
		"code": code,
		"msg": msg,
	})
}