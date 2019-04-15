package api

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/ddliu/go-httpclient"
)

//Init add basic confi to httpclient
func Init(cookies []*http.Cookie) {
	jar, _ := cookiejar.New(nil)
	u, _ := url.Parse("https://www.zhihu.com/")
	jar.SetCookies(u, cookies)
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
		httpclient.OPT_REFERER:   "https://www.zhihu.com/",
		"Accept-Language":        "zh-CN,zh;q=0.9",
		"Accept":                 "*/*",
		httpclient.OPT_COOKIEJAR: jar,
	})
}
