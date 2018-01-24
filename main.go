package main

import (
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	go Run()
	go CheckStorage()
	for {
		time.Sleep(1000 * time.Second)
	}
}
