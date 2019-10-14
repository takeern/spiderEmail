package dao

import (
	"net/http"
)

var (
	IP_QUEUE 
)

type Dispatch struct {
	Ip_list		[]string
	status		bool

}

func CreateEmailDispatch(createStatus bool) *Dispatch {
	d := &Dispatch{
		status: createStatus,
		Ip_list: make([]string, 0, 100),
	}

	
}

func (d *Dispatch) newIpConnect(ip string) *Dispatch {
	for _, item := range d.Ip_list {
		if item == ip {
			Log.Warn("this ip is in route", ip)
			return d
		}
	}
	d.Ip_list = append(d.Ip_list, ip)
}

// func createServe()  {
// 	http.HandleFunc("/register", sayhelloName)
// }

// // handle 其他 ip 注册
// func handleRegister(w http.ResponseWriter, r *http.Request)  {
	
// }