package master

import (
	"github.com/google/martian/log"
	"sync"
	"fmt"
	"spider/interval/conf"
	"math/rand"
	"context"
	"time"
	pb "spider/interval/serve/grpc"
	"spider/interval/dao/utils"
	"google.golang.org/grpc"
)

var (
	wg 		*sync.WaitGroup
)

type EmailDispatch struct {
	mu      	sync.Mutex // guards balance
	c			pb.TaskClient
	Ip_list		[]string
	Email_list 	[]string
	Error_Email_list	[]string
	Success_Email_list	[]string
	Email_send_index	int
	send_user_index		int
	modalDb		*utils.ModalDb
}	

func CreateEmailDispatch(url string) *EmailDispatch {
	d := &EmailDispatch{
		Ip_list: make([]string, 0, 100),
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

func (d *EmailDispatch) newIpConnect(ip string) {
	for _, item := range d.Ip_list {
		if item == ip {
			utils.Log.Info("this ip is in route", ip)
			return 
		}
	}
	d.Ip_list = append(d.Ip_list, ip)
	d.c, _ = CreateConn(ip)
}



func (d *EmailDispatch) InjectIp(ips []string) {
	d.Ip_list = ips
	return
}

func (d *EmailDispatch) InjectUrl(url string) {
	d.modalDb.Close()
	d.modalDb = utils.NewDb(url)
	list, err := d.modalDb.SelectData(1000)
	d.Email_list = make([]string, 0, 3000)
	for _, item := range list {
		d.Email_list = append(d.Email_list, item.Email)
	}
	if err != nil {
		utils.Log.Warn("get emails error", url, err)
	} else {
		utils.Log.Info("get emails success", len(d.Email_list))
	}
	return
}

func (d *EmailDispatch) HandleNewIpRegistry(ip string) {
	fmt.Println(ip)
	go sendTask(ip, d.send_user_index, d)
	// go func ()  {
	// 	fmt.Println("test")
	// 	for {
	// 		for a := 0; a < 10; a++ {
	// 			continue
	// 			d.mu.Lock()
	// 			fmt.Printf("a 的值为: %d\n", a)
	// 			d.mu.Unlock()
	// 		}
	// 		time.Sleep((conf.WAIT_SEND_EMAIL_TIME + time.Duration(rand.Intn(3 * 60 * 60))) * time.Second)
	// 	}
	// }()
}

/*
 * ip 当前连接地址
 * index 当前使用第几套send list 模型
 */
func sendTask(ip string, index int, d *EmailDispatch) {
	c, err := CreateConn(ip)
	Aclen := len(conf.SendList[0])
	fmt.Println(Aclen)
	var sendModalIndex [4]int
	if err != nil {
		fmt.Println(err)
		utils.Log.Error("connect error", ip, err)
		return
	}
	if (index > Aclen - 1) {
		utils.Log.Error("outside sendlist", index)
		return
	}
	acList := conf.SendList[index]
	for {
		for i, item := range acList {
			
			if (d.Email_send_index > len(d.Email_list) - 1) {
				utils.Log.Error("email send complete", d.Email_send_index)
				continue
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
			fmt.Println(resp)
			d.Email_send_index++
			if err != nil || resp.Code != 10000 {
				d.Error_Email_list = append(d.Error_Email_list, email)
				utils.Log.Error("grpc: send email error %v, ac: %s, ", &resp.ErrorMsg, item.Ac, email, i)
				d.mu.Unlock()
				continue
			} else {
				// 执行成功
				utils.Log.Info("send email: success", item.Ac, email, i)
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


func CreateConn(ip string) (pb.TaskClient, error) {
	
	conn, err := grpc.Dial(ip + ":" + conf.SLAVE_PORT, grpc.WithInsecure())
	if err != nil {
		log.Errorf("connet to slave node failed, node ip: %s, err: %v", ip, err)
		return nil ,err
	}
	c := pb.NewTaskClient(conn)
	return c, err
}
// func createServe()  {
// 	http.HandleFunc("/register", sayhelloName)
// }

// // handle 其他 ip 注册
// func handleRegister(w http.ResponseWriter, r *http.Request)  {
	
// }