package model

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

const (
	HttpProxyList  string = "http_proxy_list"
	HttpsProxyList string = "https_proxy_list"

	CopyHttpProxyList  string = "copy_http_proxy_list"
	CopyHttpsProxyList string = "copy_https_proxy_list"
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:               "127.0.0.1:6379",
		Password:           "",
		DB:                 0,
		MaxRetries:         2,
		PoolSize:           20,
		IdleTimeout:        60 * time.Second,
		IdleCheckFrequency: 5 * time.Second,
	})
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func SaveOne(ipinfo *IpInfo) error {
	if ipinfo.PxyType == PXY_TYPE_HTTP {
		return SaveOneHttp(ipinfo)
	} else {
		return SaveOneHttps(ipinfo)
	}
}

func SaveOneHttp(ipinfo *IpInfo) error {
	ipJson, e := json.Marshal(ipinfo)
	if e != nil {
		return e
	}

	isExist := redisClient.SIsMember(HttpProxyList, string(ipJson)).Val()
	if isExist == true {
		return nil
	}
	err := redisClient.SAdd(HttpProxyList, string(ipJson)).Err()
	return err
}

func SaveOneHttps(ipinfo *IpInfo) error {
	ipJson, e := json.Marshal(ipinfo)
	if e != nil {
		return e
	}
	isExist := redisClient.SIsMember(HttpsProxyList, string(ipJson)).Val()
	if isExist == true {
		return nil
	}
	err := redisClient.SAdd(HttpsProxyList, string(ipJson)).Err()
	return err
}

func DecodeOne(ipJson string) (*IpInfo, error) {
	var ipInfo *IpInfo
	err := json.Unmarshal([]byte(ipJson), &ipInfo)
	if err != nil {
		return &IpInfo{}, err
	}

	return ipInfo, nil
}
