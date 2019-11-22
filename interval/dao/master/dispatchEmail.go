package master

import (
	"sync"
	"spider/interval/conf"
	"math/rand"
	"context"
	"time"
	pb "spider/interval/serve/grpc"
	"spider/interval/dao/utils"
	"spider/interval/modal"
)

var (
	wg 		*sync.WaitGroup
)

type EmailDispatch struct {
	mu      				sync.Mutex // guards balance
	c						pb.TaskClient
	Ip_list					*modal.Queue
	Close_ip_list			*modal.Queue
	Email_list 				[]string
	Error_Email_list		[]string
	Success_Email_list		[]string
	Email_send_index		int
	send_user_index			int
	modalDb					*utils.ModalDb
}	

func CreateEmailDispatch(url string) *EmailDispatch {
	d := &EmailDispatch{
		Ip_list: modal.NewQueue(100),
		Close_ip_list: modal.NewQueue(100),
		Email_list: make([]string, 0, 3000),
		Error_Email_list: make([]string, 0, 3000),
		Success_Email_list: make([]string, 0, 3000),
		modalDb: utils.NewDb(url),
	}
	// list, err := d.modalDb.SelectData(1000)
	// if err != nil {
	// 	utils.Log.Warn("get emails error", url, err)
	// } else {
	// 	utils.Log.Info("get emails success", len(d.Email_list))
	// }
	// for _, item := range list {
	// 	d.Email_list = append(d.Email_list, item.Email)
	// }

	for i := 0; i < 15; i++ {
		d.Email_list = append(d.Email_list, "takeern@163.com")
	}

	return d
}

func (d *EmailDispatch) HandleNewIpRegistry(ip string, c pb.TaskClient) (code int, msg string) {
	if (!d.Ip_list.HasValue(ip)) {
		d.Ip_list.Push(ip)
		go sendTask(ip, d.send_user_index, d, c)
		code = conf.RegisterCodeSuccess
		msg = conf.RegisterMsgSuccess
	} else {
		code = conf.RegisterCodeError
		msg = conf.RegisterMsgErrorRepeat
	}

	return code, ip + msg + "task: send email"
}

func (d *EmailDispatch) closeIp(ip string) {
	d.Ip_list.Remove(ip)
	d.Close_ip_list.Push(ip)
	utils.Log.Info("remove connect ip: " + ip + "task: email")
}

/*
 * ip 当前连接地址
 * index 当前使用第几套send list 模型
 */
func sendTask(ip string, index int, d *EmailDispatch, c pb.TaskClient) {
	var sendModalIndex [4]int
	var error_spider_times int
	Aclen := len(conf.SendList[0])
	acList := conf.SendList[index]

	if (index > Aclen - 1) {
		d.closeIp(ip)
		utils.Log.Error("outside sendlist", index)
		return
	}

	for {
		for i, item := range acList {
			
			if (d.Email_send_index > len(d.Email_list) - 1) {
				utils.Log.Error("email send complete", d.Email_send_index)
				d.closeIp(ip)
				return
			}

			d.mu.Lock()
			email := d.Email_list[d.Email_send_index]
			req := &pb.HandleTaskReq{
				TaskCode: 1000,
				EmailInfo: &pb.EmailInfo{
					Ac: item.Ac,
					Ps:	item.Ps,
					Host: item.Host,
					Receive: email,
					ModalIndex: int32(sendModalIndex[i]),
				},
			}
			resp, err := c.HandleTask(context.Background(), req)
			d.Email_send_index++

			if err != nil || resp.Code != 10000 {
				d.Error_Email_list = append(d.Error_Email_list, email)
				utils.Log.Error("grpc: send email error %v, ac: %s, ", &resp.ErrorMsg, item.Ac, email, i)
				d.mu.Unlock()
				error_spider_times ++
				continue
			} else {
				// 执行成功
				utils.Log.Info("send email: success", item.Ac, email, i)

				error_spider_times --
				if (error_spider_times < 0) {
					error_spider_times = 0
				}

				d.Success_Email_list = append(d.Success_Email_list, email)
				d.modalDb.UpdateStatus(email, true)
			}
			d.mu.Unlock()

			sendModalIndex[i] ++
			if (sendModalIndex[i] > len(conf.EmailModalList) - 1) {
				sendModalIndex[i] = 0
			}
		}
		time.Sleep((conf.WAIT_SEND_EMAIL_TIME + time.Duration(rand.Intn(3 * 60 * 60))) * time.Second)
	}
}
