package getter

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"getpxy/filter"
	"getpxy/model"

	"github.com/parnurzeal/gorequest"
)

func Ip66() (result []*model.IpInfo) {
	pollUrl := "http://www.66ip.cn/nmtq.php?getnum=300&isp=0&anonymoustype=3&start=&ports=&export=&ipaddress=&area=0&proxytype=2&api=66ip"
	_, body, err := gorequest.New().Get(pollUrl).End()

	if err != nil {
		log.Println("fail to make Request")
		return
	}
	//body = strings.TrimSpace(body)
	list := strings.Split(body, "<br />")

	//fmt.Println(str[1:300])
	list = list[1:300]
	fmt.Println(len(list))

	for _, ip := range list {
		ipArr := strings.Split(strings.TrimSpace(ip), ":")

		if ipArr[0] == "" || ipArr[1] == "" {
			continue
		}

		ipinfo := &model.IpInfo{
			Ip:      ipArr[0],
			Port:    ipArr[1],
			PxyType: model.PXY_TYPE_HTTP,
			Level:   model.ANONYMITY_LEVEL_GENERAL,
		}
		result = append(result, ipinfo)

	}
	return
}

func Get66Ip() {
	allList := Ip66()
	getCount := len(allList)
	log.Println("从66ip网页共获取数据：", getCount)

	var wg sync.WaitGroup
	for _, ip := range allList {
		wg.Add(1)
		go func(ip *model.IpInfo) {
			if filter.CheckIP(ip) == true {
				log.Println("Available 66ip ip: ", ip)
				result = append(result, ip)
			}
			wg.Done()
		}(ip)
	}
	wg.Wait()
	log.Println("过滤66ip网页数据后，可用数据：", len(result))
	return

}
