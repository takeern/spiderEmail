package master

import (
	"sync"
	"time"
	"crypto/tls"
	"regexp"
	"net/http"
	"strings"
	"context"
	"fmt"

	pb "spider/interval/serve/grpc"
	"spider/interval/dao/utils"
	// "google.golang.org/grpc"
	"spider/interval/conf"
	"spider/interval/modal"
)

type SpiderDispatch struct {
	mu      				sync.Mutex
	c						pb.TaskClient
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
    modalDb 				*utils.ModalDb
}

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func CreateDispatchSpider(url string) *SpiderDispatch {
	host_url := getHostUrl(url)
	d := &SpiderDispatch{
		Ip_list: modal.NewQueue(100),
		Close_ip_list: modal.NewQueue(100),
		Wait_spider_queue: modal.NewQueue(2000),
		Had_spider_queue: modal.NewQueue(2000),
		Error_spider_queue: modal.NewQueue(2000),
		Host_url: host_url,
		modalDb: utils.NewDb(host_url),
		Cache_email: make(map[string]string),
	}
	d.AppendUrl(url)
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

func (d *SpiderDispatch)AppendUrl(url string) {
	d.Wait_spider_queue.Push(url)
}

func (d *SpiderDispatch) HandleNewIpRegistry(ip string) (code int, msg string){
	if (!d.Ip_list.HasValue(ip)) {
		d.Ip_list.Push(ip)
		go d.sendTask(ip)
		code = conf.RegisterCodeSuccess
		msg = conf.RegisterMsgSuccess
	} else {
		code = conf.RegisterCodeError
		msg = conf.RegisterMsgErrorRepeat
	}

	return code, ip + msg + "task: spider"
}

func (d *SpiderDispatch) closeIp(ip string) {
	d.Ip_list.Remove(ip)
	d.Close_ip_list.Push(ip)
	utils.Log.Info("remove connect ip: " + ip + "task: spider")
}

func (d *SpiderDispatch) sendTask(ip string) {
	c, err := CreateConn(ip)
	var error_spider_times int
	if err != nil {
		utils.Log.Error("connect grpc error", ip, err)
		return
	} else {
		utils.Log.Info("connect grpc success", ip)
	}

	for {
		if (error_spider_times > conf.Retry_Spider_Times) {
			d.closeIp(ip)
			return
		}
		d.mu.Lock()
		d.Spider_times ++
		if (d.Wait_spider_queue.Len() != 0) {
			next_url := d.Wait_spider_queue.Shift()
			utils.Log.Info("next Url", next_url)
			req := &pb.HandleTaskReq{
				TaskCode: conf.SPIDER_EMAIL,
				SpiderUrl: next_url,
			}

			resp, err := c.HandleTask(context.Background(), req)
			if err != nil || resp.Code != 10000 {
				d.Error_spider_queue.Push(next_url)
				msg := ""
				if resp != nil {
					msg = resp.ErrorMsg
				}
				fmt.Println(resp)
				error_spider_times ++
				utils.Log.Error("grpc: spider url error ", next_url, msg)
			} else {
				// 爬取成功
				error_spider_times --
				if (error_spider_times < 0) {
					error_spider_times = 0
				}
				d.Had_spider_queue.Push(next_url)
				utils.Log.Info("spider url: success", next_url)

				for _, url := range resp.SpiderInfo.Urls {
					if (!d.Wait_spider_queue.HasValue(url) && !d.Had_spider_queue.HasValue(url)) {	// 检查是否已爬去过
						d.Wait_spider_queue.Push(url)
					}
				}
				for _, email := range resp.SpiderInfo.Emails {
					_, ok := d.Cache_email[email]
					if !ok {
						d.modalDb.InsertData(next_url, email)
						d.Cache_email[email] = next_url
					}
				}
			}
			d.mu.Unlock()
		} else {
			d.closeIp(ip)
			utils.Log.Info("spider over")
			return
		}

		time.Sleep(conf.WAIT_SPIDER_TIME * time.Second)
	}
}
