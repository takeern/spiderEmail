package dao

import (
    "net/smtp"
	"strings"
	"spider/interval/conf"
	"time"
	"fmt"
	"sync"
)

//发送邮件的逻辑函数
func SendMail(user, password, host, to, subject, body, mailtype string) error {
    hp := strings.Split(host, ":")
    auth := smtp.PlainAuth("", user, password, hp[0])
    var content_type string
    if mailtype == "html" {
        content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
    } else {
        content_type = "Content-Type: text/plain" + "; charset=UTF-8"
    }

    msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
    send_to := strings.Split(to, ";")
    err := smtp.SendMail(host, auth, user, send_to, msg)
    return err
}

func userSend(ac, ps, host string) {
	fmt.Println("send")
	var modalIndex, rcIndex int
	emLen := len(conf.EmailModalList)
	rcLen := len(conf.RecieveList)
	// ticker := time.NewTicker( conf.WAIT_SEND_EMAIL_TIME * time.Second)
	// defer ticker.Stop()
	for {
		fmt.Println("send")
		if (modalIndex > emLen ) {
			modalIndex = 0
		}

		if (rcIndex > rcLen) {
			rcIndex = 0
		}

		err := SendMail(ac, ps, host, conf.RecieveList[rcIndex], "无主题", conf.EmailModalList[modalIndex], "html")
		if err != nil { 
			Log.Warn("send email error: ", err, ac, conf.RecieveList[rcIndex])
		} else {
			Log.Info("发送邮件成功", ac, conf.RecieveList[rcIndex])
		}
		time.Sleep(conf.WAIT_SEND_EMAIL_TIME * time.Second)
	}
	if (modalIndex > emLen ) {
		modalIndex = 0
	}

	if (rcIndex > rcLen) {
		rcIndex = 0
	}

	err := SendMail(ac, ps, host, conf.RecieveList[rcIndex], "无主题", conf.EmailModalList[modalIndex], "html")
	if err != nil { 
		Log.Warn("send email error: ", err, ac, conf.RecieveList[rcIndex])
	} else {
		Log.Info("发送邮件成功", ac, conf.RecieveList[rcIndex])
	}
	time.Sleep(conf.WAIT_SEND_EMAIL_TIME * time.Second)
	// for {
	// 	select {
	// 	//此处在等待channel中的信号，因此执行此段代码时会阻塞120秒
	// 	case <-ticker.C:

	// 		if (modalIndex > emLen ) {
	// 			modalIndex = 0
	// 		}

	// 		if (rcIndex > rcLen) {
	// 			rcIndex = 0
	// 		}

	// 		err := SendMail(ac, ps, host, conf.RecieveList[rcIndex], "无主题", conf.EmailModalList[modalIndex], "html")
	// 		if err != nil { 
	// 			Log.Warn("send email error: ", err, ac, conf.RecieveList[rcIndex])
	// 		} else {
	// 			Log.Info("发送邮件成功", ac, conf.RecieveList[rcIndex])
	// 		}
	// 	}
	// }
}

func GoSend() {
	var wg = sync.WaitGroup{}
	wg.Add(len(conf.SendList))
	for _, user := range conf.SendList {
		go func () {
			userSend(user.Ac, user.Ps, user.Host)
			wg.Done()
		}()
	}
	wg.Wait()
}
