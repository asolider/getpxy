package main

import (
	"fmt"
	"net/http"
	"os"
	//"compress/gzip"

	goquery "github.com/PuerkitoBio/goquery"
)

var putFile = "/www/mycode/html.text"

func main() {
	fmt.Println("hello")
	getHtml()
}

func getHtml() {
	request, err := http.NewRequest("GET", "http://www.xicidaili.com/nn", nil)
	if err != nil {
		fmt.Println("fail to create request")
		os.Exit(1)
	}

	request.Header.Add("User-Agent", getUserAgent())
	request.Header.Add("Accept", getAccept())
	request.Header.Add("Accept-Language", getAcceptLanguage())
	request.Header.Add("Accept-Encoding", getAcceptEncoding())
	request.Header.Add("Cache-Control", getCacheControl())
	request.Header.Add("Connection", getConnection())

	httpClient := &http.Client{}

	response, err := httpClient.Do(request)

	if err != nil || response.StatusCode != 200 {
		fmt.Println("fail to get content")
		os.Exit(2)
	}
	defer response.Body.Close()
	//f, err := os.OpenFile(putFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//
	//buf := make([]byte, 1024)
	//unzipbody, _ := gzip.NewReader(response.Body)
	//for {
	//	n, _ := unzipbody.Read(buf)
	//	if 0 == n {
	//		break
	//	}
	//	f.WriteString(string(buf[:n]))
	//}

	doc, e := goquery.NewDocumentFromResponse(response)
	if e != nil {
		fmt.Println("read doc err")
		os.Exit(3)
	}

	doc.Find("#ip_list tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		ip := s.Find("td").Eq(1).Text()
		port := s.Find("td").Eq(2).Text()
		httptype := s.Find("td").Eq(5).Text()
		fmt.Println(ip + "  " +port+ "  "+httptype)
	})


}
