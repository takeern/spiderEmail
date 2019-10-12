package dao

import (
	"time"
	"net/smtp"
	"io/ioutil"
	"crypto/tls"
	"regexp"
	"net/http"
	"strings"

	"spider/interval/modal"
	pb "spider/interval/serve/grpc"
)

type Server struct{}

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client = http.Client{
		Timeout: modal.SPIDER_TIMEOUT * time.Second,
	}
}

func (s *Server) HandleTask(ctx context.Context, req *pb.HandleTaskReq) (*pb.HandleTaskResp) {
	resp := new(pb.HandleTaskResp)
	selectTaskRun(req, resp)
	return resp
}

/*
 * 选择任务类型
 * 1000 发送 email
 * 1001 爬取页面 email
 */
func selectTaskRun(req *pb.HandleTaskReq, resp *pb.HandleTaskResp) (*pb.HandleTaskResp) {
	switch req.TaskCode {
	case conf.SEND_EMAIL:
		err := sendMail(req.EmailInfo.Ac, 
			req.EmailInfo.Ps, 
			req.EmailInfo.Host, 
			req.EmailInfo.Receive,
			"无主题",
			conf.EmailModalList[req.EmailInfo.ModalIndex],
			"html")
		if err != nil {
			resp.Code = 10001
			resp.ErrorMsg = err.Error()
		} else {
			resp.Code = 10000
		}
		
		return resp
	case conf.SPIDER_EMAIL:
		err, emails, urls := spiderEmail(req.SpiderUrl, 0)
		if err != nil {
			resp.Code = 10002
			resp.ErrorMsg = err.Error()
		} else {
			resp.Code = 10000
			resp.SpiderInfo.Emails = emails
			resp.SpiderInfo.Urls = urls
		}
		return resp
	default:
		resp.Code = 10003
		resp.ErrorMsg = "unhandle task code"
		return resp
	}
}

// 发送邮件任务
func sendMail(user, password, host, to, subject, body, mailtype string) error {
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

// 爬取页面任务
func spiderEmail(url string, times int) (error, []string, []string) {
	emails, urls := make([]string, 0, 100)
	if (times > modal.HTTP_TRY_REQUEST_TIMES) {
		return new Error("too many try"), emails, urls
	}
	res, err := client.Get(url)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("warn: http err", err.Error())
		times ++
		return spiderEmail(url, times)
	} else {
		Body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err);
		}
		html := string(Body)
		emails = drawEmail(html)
		urls = drawUrl(html)
		return nil, emails, urls
	}
}

// 提取页面邮箱
func drawEmail(html string) []string {
	re := regexp.MustCompile(`[a-zA-Z0-9_\-\.]+@[a-zA-Z0-9]+\.[a-zA-Z0-9\.]+`)
	params := re.FindAllSubmatch([]byte(html), -1)
	emails := make([]string, 0, 100)
	for _, param := range params {
		emails = append(emails, string(param[0]))
	}
	return emails
}

// 提取页面url
func drawUrl(html string) {
	re := regexp.MustCompile(`<a[^>]*href[=\"\'\s]+([^\"\']*)[\"\']?[^>]*>`)
	params := re.FindAllSubmatch([]byte(html), -1)
	urls := make([]string, 0, 100)
	for _, param := range params {
		url := editUlr(string(param[1]))
		urls = append(urls, url)
	}
	return urls
}

// 检查 url 合法性
func editUlr(url string) (string) {
	isAbsoluteUrl, ok := regexp.MatchString(`(http|https):\/\/`, url)
	if ok != nil {
		fmt.Println("isAbsoluteUrl error", url, ok)
		return ""
	}
	if (isAbsoluteUrl) {
		iscors, ok := regexp.MatchString(host_url, url)
		if ok != nil {
			fmt.Println("iscors error", url, ok)
			return ""
		}
		if (iscors) {
			return url
		}
		return ""
	} else {
		ok, _ := regexp.MatchString(`javascript`, url)
		if (ok) {
			return ""
		} else {
			return host_url + url
		}
	}
}