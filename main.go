package main

import (
	"getpxy/getter"
	"time"
)

func main() {
	//	go Run()
	//	go CheckStorage()
	getter.Ip66()
	for {
		time.Sleep(1000 * time.Second)
	}
}
