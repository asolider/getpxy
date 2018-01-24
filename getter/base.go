package getter

import (
	"getpxy/filter"
	"getpxy/model"
	"log"
	"sync"
)

type GetData struct {
	SourceName string
	SourceUrl  string

	GetSource func(*GetData) []*model.IpInfo

	// 抓取的原始数据
	OriginData []*model.IpInfo
	// 过滤后的数据
	FilterData []*model.IpInfo
}

func (this *GetData) FilterSource() {
	originLen := len(this.OriginData)
	log.Printf("抓取 %s 数据共 %d条", this.SourceName, originLen)

	if originLen == 0 {
		log.Printf("本次 %s 无数据，结束过滤", this.SourceName)
	}

	var wg sync.WaitGroup
	for _, ip := range this.OriginData {
		wg.Add(1)
		go func(ip *model.IpInfo) {
			if filter.CheckIP(ip) == true {
				log.Println("AvailableXici ip: ", ip)
				this.FilterData = append(this.FilterData, ip)
			}
			wg.Done()
		}(ip)
	}
	wg.Wait()
	log.Printf("过滤 %d 网页数据后，可用数据：", len(this.FilterData))
	return
}

func (this *GetData) Save() {
	filterLen := len(this.FilterData)
	log.Printf("抓取可用 %s 数据共 %d条", this.SourceName, filterLen)

	if filterLen == 0 {
		log.Printf("抓取 %s 无可用数据，无需保存 ", this.SourceName)
	}

	for _, ip := range this.FilterData {
		model.SaveOne(ip)
	}
	log.Printf("%s 保存完毕", this.SourceName)
}

func (this *GetData) Run() {
	this.OriginData = this.GetSource(this)
	this.FilterSource()
	this.Save()
	log.Printf("%s 本次执行完毕", this.SourceName)
}
