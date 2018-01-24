package getter

import (
	"log"
	"net/http"
	"strings"

	"getpxy/model"
	"getpxy/util"

	"github.com/PuerkitoBio/goquery"
)

func XiciGet(getData *GetData) (originData []*model.IpInfo) {
	request, err := http.NewRequest("GET", getData.SourceUrl, nil)
	if err != nil {
		log.Printf("%s fail to NewRequest", getData.SourceName)
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
		log.Printf("fail to get [%s] content", getData.SourceName)
		return
	}
	defer response.Body.Close()

	doc, e := goquery.NewDocumentFromResponse(response)
	if e != nil {
		log.Println("fail to analyze [%s] content", getData.SourceName)
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
			originData = append(originData, ipinfo)
		}
	})
	return
}

var XiciDaili = &GetData{
	SourceName: "xicidaili",
	SourceUrl:  "http://www.xicidaili.com/nn/1",
	GetSource:  XiciGet,
}
