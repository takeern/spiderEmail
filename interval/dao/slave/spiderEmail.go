package slave

import (
	"time"
	"io/ioutil"
	"crypto/tls"
	"regexp"
	"net/http"
	"errors"
	"spider/interval/conf"
	"code.sajari.com/docconv"
)

var (
	client http.Client
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client = http.Client{
		Timeout: conf.SPIDER_TIMEOUT * time.Second,
	}
}

// 爬取页面任务
func SpiderEmail(url string, times int) (error, []string, []string) {
	emails := make([]string, 0, 100)
	urls := make([]string, 0, 100)
	var html string
	if (times > conf.HTTP_TRY_REQUEST_TIMES) {
		return errors.New("too many try"), emails, urls
	}
	res, err := client.Get(url)
	if err != nil {
		times ++
		return SpiderEmail(url, times)
	} else {
		defer res.Body.Close()
		isPDF, _ := regexp.MatchString(`pdf`, res.Header["Content-Type"][0])
		if isPDF {
			html, _, _ = docconv.ConvertPDF(res.Body)
		} else {
			Body, _ := ioutil.ReadAll(res.Body)
			html = string(Body)
		}
	}

	emails = drawEmail(html)
	re := regexp.MustCompile(`(http|https):\/\/?([^/]*)`)
	host_url := string(re.Find([]byte(url)))
	urls = drawUrl(html, host_url)
	return err, emails, urls
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
func drawUrl(html string, host_url string) []string {
	reScrpit := regexp.MustCompile(`<script[^>]*?>(?:.|\n)*?<\/script>`)
	html = string(reScrpit.ReplaceAll([]byte(html), []byte("")))
	re := regexp.MustCompile(`<a[^>]*href[=\"\'\s]+([^\"\']*)[\"\']?[^>]*>`)
	params := re.FindAllSubmatch([]byte(html), -1)
	urls := make([]string, 0, 100)
	for _, param := range params {
		url := editUlr(string(param[1]), host_url)
		urls = append(urls, url)
	}
	return urls
}

// 检查 url 合法性
func editUlr(url string, host_url string) (string) {
	isAbsoluteUrl, ok := regexp.MatchString(`(http|https):\/\/`, url)
	if ok != nil {
		return ""
	}
	if (isAbsoluteUrl) {
		iscors, ok := regexp.MatchString(host_url, url)
		if ok != nil {
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
