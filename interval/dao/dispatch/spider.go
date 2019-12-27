package dispatch

import (
	"time"
	"regexp"
	"strings"
	"context"
	"sort"
	// "fmt"

	pb "spider/interval/serve/grpc"
	"spider/interval/dao/utils"
	// "spider/interval/dao/master"
	"spider/interval/conf"
	"spider/interval/modal"
)

type SpiderData struct {
	Wait_spider_queue 		*modal.Queue
	Had_spider_queue 		*modal.Queue
	retry_spider_queue 		*modal.Queue
	Error_spider_queue 		*modal.Queue
	Cache_email 			map[string]string
	Host_url 				string
	Spider_times 			int
}

type IState struct {
	sleep					bool
	ch						chan *ICh
	ctx						context.Context
	cancel					context.CancelFunc
	ipList				    map[string]bool
	conns					map[string]pb.TaskClient
}

type Spider struct {
	Data					*SpiderData
	recordData				[]*pb.SpiderRecordData
	state					*IState
	modalDb 				*utils.ModalDb
	ipState					map[string]int
}

type ICh struct{
	req 					*pb.HandleTaskReq
	connC					*IConnC
}

type IConnC struct {
	c						pb.TaskClient
	ip 						string
}

/*	生命周期 dispatch 
 *	1. new
 *  2. initData		初始化数据
 *	3. run			开始执行任务
 *  4. exit			退出 gc
 */

// 创建 spider 
func NewSpider (ctx	context.Context, url string, sleep bool, ipList map[string]bool, conns map[string]pb.TaskClient) *Spider {
	ctxD, cancel := context.WithCancel(ctx)
	d := &Spider {
		recordData: make([]*pb.SpiderRecordData, 0, 1000),
		ipState: make(map[string]int),
		state: &IState{
			sleep: sleep,
			ch: make(chan *ICh, 0),
			cancel: cancel,
			ctx: ctxD,
			ipList: ipList,
			conns: conns,
		},
	}

	if !sleep {
		host_url := getHostUrl(url)
		d.initData(host_url)
		d.modalDb = utils.NewDb(host_url)
		d.Data.Wait_spider_queue.Push(url)
		d.Run()
	}
	return d
}

// 修饰 url 成为 db 中 唯一 key
func getHostUrl(url string) string {
	re := regexp.MustCompile(`(http|https):\/\/?([^/]*)`)
	url = string(re.Find([]byte(url)))
	hasPrex, _ := regexp.MatchString(`[\/]$`, url)
	if (!hasPrex) {
		url += "/"
	}
	return strings.Replace(url, ".", "", -1)
}

//  初始化数据
func (d *Spider) initData(host_url  string) {
	d.Data = &SpiderData {
		Wait_spider_queue: modal.NewQueue(2000),
		Had_spider_queue: modal.NewQueue(2000),
		Error_spider_queue: modal.NewQueue(2000),
		Host_url: host_url,
		Cache_email: make(map[string]string),
	}
}

// Run
func (d *Spider) Run() {
	go d.dispatch(d.state.ctx, d.state.ch)
	go d.handleSend(d.state.ctx, d.state.ch)
}

// 退出 所有该 dispatch routine
func (d *Spider) Exit() {
	d.state.cancel()
}

// 分发任务
func (d *Spider) dispatch(ctx context.Context, out chan<- *ICh) {
	for {
		select {
		case <- ctx.Done():
			utils.Log.Info(" 分发任务 goroutine exit by single, url:", d.Data.Host_url)
			return
		default:
			// 正常发送流程
			conns := d.getConns()
			n := len(conns)
			utils.Log.Info("开始分发任务， ip 数量: ", len(conns))
			if n == 0 {
				time.Sleep(conf.WAIT_SPIDER_TIME * time.Second)
			} else {
				for _, v := range conns {
					if (d.Data.Wait_spider_queue.Len() == 0) {
						utils.Log.Info(" 分发任务 goroutine exit by spider all:", d.Data.Host_url)
						return
					}
					d.Data.Spider_times ++
					targetUrl := d.Data.Wait_spider_queue.Shift()
					msg := &ICh{
						req: &pb.HandleTaskReq{
							TaskCode: conf.SPIDER_EMAIL,
							SpiderUrl: targetUrl,
						},
						connC: v,
					}
					out <- msg
					time.Sleep(time.Duration(conf.WAIT_SPIDER_TIME / n) * time.Second) // 睡眠时间
				}
			}
		}
	}
}

// 发送 task
func (d *Spider) handleSend(ctx context.Context, in <-chan *ICh) {
	for {
		select {
		case <-ctx.Done():
			utils.Log.Info(" 发送 task goroutine exit, url:", d.Data.Host_url)
			return
		case msg := <-in:
			utils.Log.Info(" 发送 spider req, targetUrl:" + msg.req.SpiderUrl + "ip :" + msg.connC.ip)
			resp, err := msg.connC.c.HandleTask(context.Background(), msg.req)
			d.recordData = append(d.recordData, &pb.SpiderRecordData{
				TargetUrl: msg.req.SpiderUrl,
				Resp: resp,
				Ip: msg.connC.ip,
			})
			d.handleResp(msg.req.SpiderUrl, resp, err, msg.connC.ip)
		}
	}
}

// 处理 grpc resp
// 将 slave 节点给的 爬取数据, 存储到内存中
func (d *Spider) handleResp(targetUrl string, resp *pb.HandleTaskResp, err error, ip string) {
	if err != nil {
		utils.Log.Error("grpc 连接失败 ip: ", ip)
		d.Data.Wait_spider_queue.Push(targetUrl)
		d.ipState[ip] ++
		return
	}
	if resp.Code != 10000 {
		d.Data.Error_spider_queue.Push(targetUrl)
		msg := ""
		if resp != nil {
			msg = resp.ErrorMsg
		}
		d.ipState[ip] ++
		utils.Log.Error("grpc: 爬取页面 error ", targetUrl, msg)
	} else {
		// 爬取成功
		d.ipState[ip] --
		if (d.ipState[ip] < 0) {
			d.ipState[ip] = 0
		}
		d.Data.Had_spider_queue.Push(targetUrl)
		utils.Log.Info("grpc: spider url success", targetUrl)

		for _, url := range resp.SpiderInfo.Urls {
			if (!d.Data.Wait_spider_queue.HasValue(url) && !d.Data.Had_spider_queue.HasValue(url) && !d.Data.Error_spider_queue.HasValue(url)) {	// 检查是否已爬去过
				d.Data.Wait_spider_queue.Push(url)
			}
		}
		for _, email := range resp.SpiderInfo.Emails {
			_, ok := d.Data.Cache_email[email]
			if !ok {
				if !d.state.sleep {
					d.modalDb.InsertData(targetUrl, email)
				}
				d.Data.Cache_email[email] = targetUrl
			}
		}
	}
}

// 获取可用句柄
func (d *Spider) getConns () []*IConnC {
	allConns := d.state.conns
	ip := d.state.ipList
	list := make([]string, 0, len(ip))
	conns := make([]*IConnC, 0, len(allConns))
	for k, ok := range ip {
		if ok {
			list = append(list, k)
		}
	}
	sort.Strings(list)

	for _, ip := range list {
		utils.Log.Info("ip: " + ip + " 状态: ", d.ipState[ip])
		if d.ipState[ip] > conf.Retry_Spider_Times{	// 该ip 错误过多 退出
			continue
		}
		grpcC, ok := allConns[ip]
		if ok {
			conns = append(conns, &IConnC{
				c: grpcC,
				ip: ip,
			})
		} else {
			utils.Log.Warn("搜索 grpc 句柄失败，存在该ip， 但是找不到grpc client", ip)
		}
	}

	return conns
}

/*
 * 同步spider dispatch数据
 * true 状态下仅同步本次时间段内 变化
 * false 情况下 获取所有需要的数据
 */
func (d *Spider) GetSyncData(record bool) (*pb.SpiderSyncData){
	r := &pb.SpiderSyncData{}
	if record {
		r.SpiderRecordData = make([]*pb.SpiderRecordData, 1, 1000)
		copy(r.SpiderRecordData, d.recordData)
		d.recordData = d.recordData[0:0]
	} else {
		r.SpiderAllData = &pb.SpiderAllData {
			WaitSpiderQueue: d.Data.Wait_spider_queue.Q,
			HadSpiderQueue: d.Data.Had_spider_queue.Q,
			ErrorSpiderQueue: d.Data.Error_spider_queue.Q,
			CacheEmail: d.Data.Cache_email,
			HostUrl: d.Data.Host_url,
		}
	}
	return r
}