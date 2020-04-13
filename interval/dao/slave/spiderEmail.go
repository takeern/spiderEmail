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
	"strings"
	"bytes"
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
	res, err := http.Get(url)
	if err != nil {
		times ++
		return SpiderEmail(url, times)
	} else {
		defer res.Body.Close()
	}


	
	Body, _ := ioutil.ReadAll(res.Body)
	html = string(Body)
	reader := bytes.NewReader(Body)
	pdf, _, _ := docconv.ConvertPDF(reader)
	pdfEmails := drawEmail(pdf)
	emails = append(drawEmail(html), pdfEmails...)

	re := regexp.MustCompile(`(http|https):\/\/?([^/]*)`)
	rePath := regexp.MustCompile(`(.*)\/`)
	path_url := string(rePath.Find([]byte(url)))
	host_url := string(re.Find([]byte(url)))
	if len(path_url) < len(host_url) {
		path_url = host_url
	}

	if !strings.HasSuffix(path_url, "/") {
		path_url = path_url + "/"
	}
	urls = drawUrl(html, host_url, path_url)
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
func drawUrl(html string, host_url string, path_url string) []string {
	reScrpit := regexp.MustCompile(`<script[^>]*?>(?:.|\n)*?<\/script>`)
	html = string(reScrpit.ReplaceAll([]byte(html), []byte("")))
	re := regexp.MustCompile(`<a[^>]*href[=\"\'\s]+([^\"\']*)[\"\']?[^>]*>`)
	params := re.FindAllSubmatch([]byte(html), -1)
	urls := make([]string, 0, 100)
	for _, param := range params {
		url := editUlr(string(param[1]), host_url, path_url)
		urls = append(urls, url)
	}
	return urls
}

// 检查 url 合法性
func editUlr(url string, host_url string, path_url string) (string) {
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
			if strings.HasPrefix(url, "/") {
				return host_url + url
			} else if strings.HasPrefix(url, "../") {
				re := regexp.MustCompile(`(.*\/)`)
				s := string(re.Find([]byte(path_url)))
				if len(s) < len(host_url) {
					s = host_url
				}
				return s + url
			} else {
				return path_url + url
			}
		}
	}
}
