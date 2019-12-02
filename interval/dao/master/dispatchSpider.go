package master

import (
	"sync"
	"time"
	"crypto/tls"
	"regexp"
	"net/http"
	"strings"
	"context"
	// "fmt"

	pb "spider/interval/serve/grpc"
	"spider/interval/dao/utils"
	"spider/interval/conf"
	"spider/interval/modal"
)

type SpiderData struct {
	Ip_list					*modal.Queue
	Close_ip_list			*modal.Queue
	Wait_spider_queue 		*modal.Queue
	Had_spider_queue 		*modal.Queue
	retry_spider_queue 		*modal.Queue
	Error_spider_queue 		*modal.Queue
	Cache_email 			map[string]string
	Current_request_url 	string
	Host_url 				string
	Spider_times 			int
}

type SpiderDispatch struct {
	mu      				sync.Mutex
	c						pb.TaskClient
	modalDb 				*utils.ModalDb
	Data					*SpiderData
	changeData				*SpiderData
	recordData				[]*pb.SpiderRecordData
	sleep					bool
}

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func initSpiderData(host_url  string) *SpiderData {
	return &SpiderData {
		Ip_list: modal.NewQueue(100),
		Close_ip_list: modal.NewQueue(100),
		Wait_spider_queue: modal.NewQueue(2000),
		Had_spider_queue: modal.NewQueue(2000),
		Error_spider_queue: modal.NewQueue(2000),
		Host_url: host_url,
		Cache_email: make(map[string]string),
	}
}

// 创建 spider 分发器
// 睡眠状态下 不开启db
func CreateSpiderDispatch(url string, sleep bool) *SpiderDispatch {
	d := &SpiderDispatch{
		sleep: sleep,
		recordData: make([]*pb.SpiderRecordData, 0, 1000),
	}
	if !sleep {
		host_url := getHostUrl(url)
		d.Data = initSpiderData(host_url)
		d.modalDb = utils.NewDb(host_url)
		d.Data.Wait_spider_queue.Push(url)
	}
	return d
}

// 修改 url 关联到表名字
func getHostUrl(url string) string {
	re := regexp.MustCompile(`(http|https):\/\/?([^/]*)`)
	url = string(re.Find([]byte(url)))
	hasPrex, _ := regexp.MatchString(`[\/]$`, url)
	if (!hasPrex) {
		url += "/"
	}
	return strings.Replace(url, ".", "", -1)
}

func (d *SpiderDispatch) HandleNewIpRegistry(ip string, c pb.TaskClient) (code int, msg string){
	if (!d.Data.Ip_list.HasValue(ip)) {
		d.Data.Ip_list.Push(ip)
		go d.sendTask(ip, c)
		code = conf.RegisterCodeSuccess
		msg = conf.RegisterMsgSuccess
	} else {
		code = conf.RegisterCodeError
		msg = conf.RegisterMsgErrorRepeat
	}

	return code, ip + msg + "task: spider"
}

// 注入初始化数据
func (d *SpiderDispatch) InjectInitData(data *pb.SpiderAllData) {
	d.Data = initSpiderData(data.HostUrl)
	d.Data.Cache_email = data.CacheEmail
	d.Data.Ip_list.PushList(data.IpList)
	d.Data.Close_ip_list.PushList(data.CloseIpList)
	d.Data.Wait_spider_queue.PushList(data.WaitSpiderQueue)
	d.Data.Had_spider_queue.PushList(data.HadSpiderQueue)
	d.Data.Error_spider_queue.PushList(data.ErrorSpiderQueue)
}

// 注入每次记录的数据
func (d *SpiderDispatch) InjectRecordData(records []*pb.SpiderRecordData)  {
	for _, item := range records {
		d.handleResp(item.TargetUrl, item.Resp, nil, 0)
		d.Data.Wait_spider_queue.Shift()
	}
}

func (d *SpiderDispatch) closeIp(ip string) {
	d.Data.Ip_list.Remove(ip)
	d.Data.Close_ip_list.Push(ip)
	utils.Log.Info("remove connect ip: " + ip + "task: spider")
}

func (d *SpiderDispatch) sendTask(ip string, c pb.TaskClient) {
	var errorTimes int

	for {
		if (errorTimes > conf.Retry_Spider_Times) {
			d.closeIp(ip)
			return
		}
		d.mu.Lock()
		d.Data.Spider_times ++
		if (d.Data.Wait_spider_queue.Len() != 0) {
			targetUrl := d.Data.Wait_spider_queue.Shift()
			utils.Log.Info("spider targetUrl", targetUrl)
			req := &pb.HandleTaskReq{
				TaskCode: conf.SPIDER_EMAIL,
				SpiderUrl: targetUrl,
			}
			d.mu.Unlock()

			resp, err := c.HandleTask(context.Background(), req)
			d.recordData = append(d.recordData, &pb.SpiderRecordData{
				TargetUrl: targetUrl,
				Resp: resp,
			})
			errorTimes = d.handleResp(targetUrl, resp, err, errorTimes)

		} else {
			d.mu.Unlock()
			d.closeIp(ip)
			utils.Log.Info("spider over")
			return
		}

		time.Sleep(conf.WAIT_SPIDER_TIME * time.Second)
	}
}

// 处理 grpc resp
// 将 slave 节点给的 爬取数据, 存储到内存中
func (d *SpiderDispatch) handleResp(targetUrl string, resp *pb.HandleTaskResp, err error, errorTimes int) (times int){
	d.mu.Lock()
	if err != nil || resp.Code != 10000 {
		d.Data.Error_spider_queue.Push(targetUrl)
		msg := ""
		if resp != nil {
			msg = resp.ErrorMsg
		}
		errorTimes ++
		utils.Log.Error("grpc: spider url error ", targetUrl, msg)
	} else {
		// 爬取成功
		errorTimes --
		if (errorTimes < 0) {
			errorTimes = 0
		}
		d.Data.Had_spider_queue.Push(targetUrl)
		utils.Log.Info("spider url: success", targetUrl)

		for _, url := range resp.SpiderInfo.Urls {
			if (!d.Data.Wait_spider_queue.HasValue(url) && !d.Data.Had_spider_queue.HasValue(url) && !d.Data.Error_spider_queue.HasValue(url)) {	// 检查是否已爬去过
				d.Data.Wait_spider_queue.Push(url)
			}
		}
		for _, email := range resp.SpiderInfo.Emails {
			_, ok := d.Data.Cache_email[email]
			if !ok {
				if !d.sleep {
					d.modalDb.InsertData(targetUrl, email)
				}
				d.Data.Cache_email[email] = targetUrl
			}
		}
	}
	d.mu.Unlock()
	return errorTimes
}

func (d *SpiderDispatch) GetAllData() (res SpiderData){
	return *d.Data
}

func (d *SpiderDispatch) GetSyncData() (res []*pb.SpiderRecordData){
	c := make([]*pb.SpiderRecordData, 1, 1000)
	copy(c, d.recordData)
	utils.Log.Info("record sync data d.recordData： ", len(d.recordData))
	d.recordData = d.recordData[0:0]
	utils.Log.Info("record sync data c： ", len(c))
	return c
}