package main

import (
	"math/rand"
	"time"
)

func getUserAgent() string {
	rand.Seed(time.Now().Unix())
	userAgent := []string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
	}
	return userAgent[rand.Intn(len(userAgent))]
}

func getAccept() string {
	return "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
}

func getAcceptLanguage() string {
	return "zh-CN,zh;q=0.9,en;q=0.8"
}

func getAcceptEncoding() string {
	return "deflate"
}

func getCacheControl() string {
	return "max-age=0"
}

func getConnection() string {
	return "keep-alive"
}