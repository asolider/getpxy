package getter

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"getpxy/model"
	"getpxy/util"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

var decodePortMap map[rune]string = map[rune]string{
	//ABCDEFGHIZ
	'A': "0", 'B': "1", 'C': "2", 'D': "3", 'E': "4", 'F': "5", 'G': "6", 'H': "7", 'I': "8", 'Z': "9",
}

func GetRealPort(port string) string {
	decodePort := ""
	for _, s := range port {
		decodePort += decodePortMap[s]
	}
	portInt, _ := strconv.Atoi(decodePort)
	return strconv.Itoa(portInt >> 3)
}

func GetRealIp(ip string) string {
    return strings.Split(ip, ":")[0]
}

func GuobanjiaGet(getData *GetData) (result []*model.IpInfo) {
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

	doc.Find("#list tbody tr").Each(func(i int, s *goquery.Selection) {
		ipPort, _ := s.Children().Eq(0).Html()
		encodePort := regexp.MustCompile("[A-Z]+").FindString(ipPort)

		ipPort = regexp.MustCompile("<pstyle=\"display:none;\">.?.?</p>").ReplaceAllString(strings.Replace(ipPort, " ", "", -1), "")
		ipPort = regexp.MustCompile("\\<[\\S\\s]+?\\>").ReplaceAllString(ipPort, "")

		levelString := strings.TrimSpace(s.Children().Eq(1).Text())
		level := model.ANONYMITY_LEVEL_GENERAL
		if levelString == "高匿" {
			level = model.ANONYMITY_LEVEL_ANVANCED
		}

		typeString := strings.TrimSpace(s.Children().Eq(2).Text())
		typeN := model.PXY_TYPE_UNKNOWN
		if typeString == "http" {
			typeN = model.PXY_TYPE_HTTP
		} else if typeString == "https" {
			typeN = model.PXY_TYPE_HTTPS
		}

		ip := &model.IpInfo{
			Ip:      GetRealIp(ipPort),
			Port:    GetRealPort(encodePort),
			PxyType: typeN,
			Level:   level,
		}
        log.Println(ip)
        result = append(result, ip)
	})
	return
}

var GuobanjiaDaili = &GetData{
	SourceName: "guobanjia",
	SourceUrl:  "http://www.goubanjia.com/free/gngn/index.shtml",
	GetSource:  GuobanjiaGet,
}
