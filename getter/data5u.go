package getter

import (
	"log"
	"strings"

	"getpxy/model"
	"getpxy/util"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

func Data5uGet(getData *GetData) (result []*model.IpInfo) {
	response, _, err := gorequest.New().Get(getData.SourceUrl).Set("User-Agent", util.GetUserAgent()).End()

	if err != nil {
		log.Printf("%s fail to NewRequest", getData.SourceName)
		return
	}
	defer response.Body.Close()

	doc, e := goquery.NewDocumentFromResponse(response)
	if e != nil {
		log.Println("fail to analyze [%s] content", getData.SourceName)
		return
	}

	doc.Find("div.wlist > ul > li").Eq(1).Children().Slice(1, -1).Each(func(i int, s *goquery.Selection) {
		pxyType := strings.TrimSpace(s.Children().Eq(3).Text())
		httpType := model.PXY_TYPE_UNKNOWN
		if strings.Contains(pxyType, "http") == true {
			httpType = model.PXY_TYPE_HTTP
		} else if strings.Contains(pxyType, "https") == true {
			httpType = model.PXY_TYPE_HTTPS
		}
		ip := &model.IpInfo{
			Ip:      strings.TrimSpace(s.Children().Eq(0).Text()),
			Port:    strings.TrimSpace(s.Children().Eq(1).Text()),
			PxyType: httpType,
			Level:   model.ANONYMITY_LEVEL_GENERAL,
		}
		result = append(result, ip)
	})
	return
}

var Data5uDaili = &GetData{
	SourceName: "data5u",
	SourceUrl:  "http://www.data5u.com/free/gngn/index.shtml",
	GetSource:  Data5uGet,
}
