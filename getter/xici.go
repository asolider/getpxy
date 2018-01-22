package getter

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"getpxy/filter"
	"getpxy/model"
	"getpxy/util"

	"github.com/PuerkitoBio/goquery"
)

func Xici() (result []*model.IpInfo) {
	pollUrl := "http://www.xicidaili.com/nn/1"

	request, err := http.NewRequest("GET", pollUrl, nil)
	if err != nil {
		log.Println("fail to NewRequest")
		return
	}

	request.Header.Add("User-Agent", util.GetUserAgent())
	request.Header.Add("Accept", util.GetAccept())
	request.Header.Add("Accept-Language", util.GetAcceptLanguage())
	request.Header.Add("Accept-Encoding", util.GetAcceptEncoding())
	request.Header.Add("Cache-Control", util.GetCacheControl())
	request.Header.Add("Connection", util.GetConnection())

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Println("fail to get [xici] content")
		return
	}
	defer response.Body.Close()

	doc, e := goquery.NewDocumentFromResponse(response)
	if e != nil {
		log.Println("fail to analyze [xici] content")
		return
	}

	doc.Find("#ip_list tr").Each(func(i int, s *goquery.Selection) {
		if i != 0 {
			ipinfo := &model.IpInfo{}

			ipinfo.Ip = strings.TrimSpace(s.Find("td").Eq(1).Text())
			ipinfo.Port = strings.TrimSpace(s.Find("td").Eq(2).Text())
			ipinfo.Level = 3

			httptype := strings.ToUpper(s.Find("td").Eq(5).Text())
			if httptype == "HTTP" {
				ipinfo.PxyType = model.PXY_TYPE_HTTP
			} else {
				ipinfo.PxyType = model.PXY_TYPE_HTTPS
			}
			result = append(result, ipinfo)
		}
	})
	return
}

// 返回可用的代理地址
func AvailableXici() (result []*model.IpInfo) {
	allList := Xici()
	log.Println("从xici网页共获取数据：", len(allList))
	if len(allList) == 0 {
		return
	}

	var wg sync.WaitGroup
	for _, ip := range allList {
		wg.Add(1)
		go func(ip *model.IpInfo) {
			if filter.CheckIP(ip) == true {
				log.Println("AvailableXici ip: ", ip)
				result = append(result, ip)
			}
			wg.Done()
		}(ip)
	}
	wg.Wait()
	log.Println("过滤xici网页数据后，可用数据：", len(result))
	return
}

func GetXici() {
	list := AvailableXici()
	if len(list) == 0 {
		return
	}
	for _, ip := range list {
		e := model.SaveOne(ip)
		log.Println(e)
	}
	log.Println("xici 执行完毕")
}
