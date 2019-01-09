package main

import (
	"fmt"
	"pipelines"
	"spiders"
	"sync"
)
import "settings"

func main() {
	fmt.Println("hello", settings.Name)
	wg := sync.WaitGroup{}
	wg.Add(2)
	pipe := make(chan string, 10)
	go func() {
		urls := make([]string, 0)
		for i:=0;i< 1000;i++{
			urls = append(urls, "https://httpbin.org/ip")
		}
		s := spiders.Httpbin{BaseSpider: spiders.BaseSpider{
			StartUrls: urls,
			MaxReq: settings.CONCURRENT_REQUESTS,
			Pipe:   pipe}}
		s.Run()
		wg.Done()
	}()
	go func (){
		p := pipelines.HttpBinPipeline{In:pipe}
		p.Process()
		wg.Done()
	}()
	wg.Wait()
}
