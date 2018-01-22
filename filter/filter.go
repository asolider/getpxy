package filter

import (
	"crypto/tls"
	"fmt"
	"getpxy/model"
	"getpxy/util"
	"time"

	"github.com/parnurzeal/gorequest"
)

func getIpProxy(ip *model.IpInfo) string {
	if ip.Ip == "" {
		return ""
	}
	pxystr := fmt.Sprintf("http://%s:%s", ip.Ip, ip.Port)

	return pxystr
}

func CheckIP(ip *model.IpInfo) bool {
	checkUrl := "http://www.baidu.com"
	ipProxy := getIpProxy(ip)
	if ipProxy == "" {
		return false
	}

	resp, _, errs := gorequest.New().
		TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetDebug(false).
		Proxy(ipProxy).
		Timeout(5*time.Second).
		Get(checkUrl).
		Set("User-Agent", util.GetUserAgent()).End()

	if errs != nil {
		return false
	}
	if resp.StatusCode == 200 {
		return true
	}
	return false
}
