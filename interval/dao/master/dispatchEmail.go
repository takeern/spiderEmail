package dao

import (
	"net/http"
	"sync"
	"spider/interval/conf"
	pb "spider/interval/serve/grpc"

	"google.golang.org/grpc"
)

var (
	mu      sync.Mutex // guards balance
)

type Dispatch struct {
	mu      	sync.Mutex // guards balance
	Ip_list		[]string
	Email_list 	[]string
	status		bool

}

func CreateEmailDispatch(createStatus bool) *Dispatch {
	d := &Dispatch{
		status: createStatus,
		Ip_list: make([]string, 0, 100),
		Email_list: make([]string, 0, 3000),
	}

	
}

func (d *Dispatch) newIpConnect(ip string) *Dispatch {
	for _, item := range d.Ip_list {
		if item == ip {
			Log.Info("this ip is in route", ip)
			return d
		}
	}
	d.Ip_list = append(d.Ip_list, ip)

}

func (d *Dispatch) InjectIp(ips []string) *Dispatch {
	d.Ip_list = ips
	return d
}

func (d *Dispatch) InjectEmail(emails []string) *Dispatch {
	d.Email_list = emails
	return d
}

func createConn(ip string) error {
	conn, err := grpc.Dial(ip + ":" + conf.SLAVE_PORT)
	if err != nil {
		log.Error("connect error", ip, err)
		return err
	}
	c := pb.NewTaskClient(conn)
}


// func createServe()  {
// 	http.HandleFunc("/register", sayhelloName)
// }

// // handle 其他 ip 注册
// func handleRegister(w http.ResponseWriter, r *http.Request)  {
	
// }