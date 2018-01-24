package getter

import (
	"log"
	"strings"

	"getpxy/model"
	"getpxy/util"

	"github.com/parnurzeal/gorequest"
)

func Ip66Get(getData *GetData) (result []*model.IpInfo) {
	_, body, err := gorequest.New().Get(getData.SourceUrl).Set("User-Agent", util.GetUserAgent()).End()

	if err != nil {
		log.Printf("%s fail to NewRequest", getData.SourceName)
		return
	}
	list := strings.Split(body, "<br />")
	list = list[1:]

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

var Ip66Daili = &GetData{
	SourceName: "66ip",
	SourceUrl:  "http://www.66ip.cn/nmtq.php?getnum=300&isp=0&anonymoustype=3&start=&ports=&export=&ipaddress=&area=0&proxytype=2&api=66ip",
	GetSource:  Ip66Get,
}
