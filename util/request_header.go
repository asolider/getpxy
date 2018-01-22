package util

import (
    "math/rand"
    "time"

)

func init() {
    rand.Seed(time.Now().UTC().UnixNano())
}

func GetUserAgent() string {
    userAgent := []string{
        "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36",
        "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0",
        "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
        "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
        "Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.9.168 Version/11.52",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
        "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_6; en-US) AppleWebKit/533.20.25 (KHTML, like Gecko) Version/5.0.4 Safari/533.20.27",
        "Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; SGH-i917)",
        "Mozilla/5.0 (compatible; WOW64; MSIE 10.0; Windows NT 6.2)",
        "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.94 Safari/537.36",
        "Opera/9.80 (Windows NT 6.1; WOW64; U; en) Presto/2.10.229 Version/11.62",
    }
    return userAgent[rand.Intn(len(userAgent))]
}

func GetAccept() string {
    return "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
}

func GetAcceptLanguage() string {
    return "zh-CN,zh;q=0.9,en;q=0.8"
}

func GetAcceptEncoding() string {
    return "deflate"
}

func GetCacheControl() string {
    return "max-age=0"
}

func GetConnection() string {
    return "keep-alive"
}
