package main

import (
	"getpxy/filter"
	"getpxy/getter"
	"getpxy/model"
	"log"
	"sync"
	"time"
)

var from []*getter.GetData

func init() {
	from = []*getter.GetData{
		getter.XiciDaili,
		getter.Ip66Daili,
	}
}

// 定时轮询各个网站抓取数据
func Run() {
	for {
		log.Println("定时抓取任务开启")
		var wg sync.WaitGroup
		for _, get := range from {
			wg.Add(1)
			go func(get *getter.GetData) {
				get.Run()
				wg.Done()
			}(get)
		}
		wg.Wait()
		log.Println("本次定时抓取任务完成，等待下一次执行")
		time.Sleep(600 * time.Second)
	}
}

// 定期检查储存的数据，剔除不可用的
func CheckStorage() {
	for {
		log.Println("定时执行清理开始...")

		CheckStorageByType(model.PXY_TYPE_HTTP)
		CheckStorageByType(model.PXY_TYPE_HTTPS)

		log.Println("定时执行清理结束，等待下一次之行")
		time.Sleep(120 * time.Second)
	}
}

func CheckStorageByType(pxyType byte) {
	redisClient := model.GetRedisClient()
	var redisSetKey string
	if pxyType == model.PXY_TYPE_HTTP {
		redisSetKey = model.HttpProxyList
	} else if pxyType == model.PXY_TYPE_HTTPS {
		redisSetKey = model.HttpsProxyList
	} else {
		log.Println("未处理的pxyTYpe", pxyType)
		return
	}

	checkList := redisClient.SMembers(redisSetKey).Val()
	log.Println(redisSetKey, " 共检查数量：", len(checkList))
	var wg sync.WaitGroup
	for _, ipJson := range checkList {
		go func(ipJson string) {
			wg.Add(1)
			ip, _ := model.DecodeOne(ipJson)
			if filter.CheckIP(ip) == false {
				log.Println("失效ip，删除...")
				redisClient.SRem(redisSetKey, ipJson)
			}
			wg.Done()
		}(ipJson)
	}
	wg.Wait()
	log.Println(redisSetKey, "检查完毕")
}
