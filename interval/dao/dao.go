package dao

import (
	"time"
	"io/ioutil"
	"crypto/tls"
	"regexp"
	"net/http"
	"fmt"
	"strings"

	"spider/interval/modal"
)



var (
	wait_spider_queue *modal.Queue
	had_spider_queue *modal.Queue
	retry_spider_queue *modal.Queue
	error_spider_queue *modal.Queue
	cache_email map[string]string
	current_request_url string
	host_url string
	spider_times int
	client http.Client
	mb *ModalDb
)


func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	cache_email	= make(map[string]string)
	wait_spider_queue = modal.NewQueue()
	had_spider_queue = modal.NewQueue()
	error_spider_queue = modal.NewQueue()
	client = http.Client{
		Timeout: modal.SPIDER_TIMEOUT * time.Second,
	}
}


func GetUrl(url string) {
	re := regexp.MustCompile(`(http|https):\/\/?([^/]*)`)
	host_url = string(re.Find([]byte(url)))
	spider_times = 3000
	hasPrex, _ := regexp.MatchString(`[\/]$`, host_url)
	if (!hasPrex) {
		host_url += "/"
	}
	fmt.Println(host_url)
	nameIndex := strings.Replace(host_url, ".", "", -1)
	mb = NewDb(nameIndex)
	parseHtml(url, 0)
}


func parseHtml(url string, times int) {
	if (times == 0) {
		had_spider_queue.Push(url)
	}

	if (times > modal.HTTP_TRY_REQUEST_TIMES) {
		fmt.Println("warn: retry max times url:%s", url)
		error_spider_queue.Push(url)
		return 
	}

	if (had_spider_queue.Len()> spider_times) {
		fmt.Println("had_spider_queue")
		return
	}
	current_request_url = url
	waitTime := time.NewTimer(time.Second * modal.SPIDER_WAIT_TIME)
	<-waitTime.C

	fmt.Println(url)

	res, err := client.Get(url)
	if err != nil {
		fmt.Println("warn: http err", err.Error())
		times ++
		parseHtml(url, times)
		return
	} else {
		defer res.Body.Close()
		Body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err);
		}
		html := string(Body)
		fmt.Println("drawEmail")
		drawEmail(html)
		drawUrl(html)
	}

	if (wait_spider_queue.Len() != 0) {
		next_url := wait_spider_queue.Shift()
		fmt.Println("爬取下一条", next_url)
		parseHtml(next_url, 0)
	}
	return
}

func drawEmail(html string) {
	re := regexp.MustCompile(`[a-zA-Z0-9_\-\.]+@[a-zA-Z0-9]+\.[a-zA-Z0-9\.]+`)
	params := re.FindAllSubmatch([]byte(html), -1)
	for _, param := range params {
		email := string(param[0])
		_, ok := cache_email[email]
		if !ok { // 此时将该 email 写入数据库
			fmt.Println(email)
			mb.InsertData(current_request_url, email)
			cache_email[email] = current_request_url
		}
	}
	// fmt.Debug(cache_email)
	fmt.Println(cache_email)
}

func drawUrl(html string) {
	re := regexp.MustCompile(`<a[^>]*href[=\"\'\s]+([^\"\']*)[\"\']?[^>]*>`)
	params := re.FindAllSubmatch([]byte(html), -1)
	for _, param := range params {
		url := editUlr(string(param[1]))
		if (len(url) != 0) { // 检查 url 符合规范
			if (!wait_spider_queue.HasValue(url) && !had_spider_queue.HasValue(url)) {	// 检查是否已爬去过
				wait_spider_queue.Push(url)
			}
		}
	}
	// fmt.Println(wait_spider_queue)
	fmt.Println(wait_spider_queue.Len())	
}


func editUlr(url string) (string) {
	// fmt.Debug("edit", url)
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
