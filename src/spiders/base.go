package spiders

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type BaseSpider struct {
	MaxReq    int
	StartUrls []string
	Pipe      chan string //输出到pipeline
}

func (s *BaseSpider) parse(body string) {
	s.Pipe <- body
}

func (s *BaseSpider) Run() {
	wg := sync.WaitGroup{}
	wg.Add(s.MaxReq)
	urls := make(chan string)
	for i := 0; i < s.MaxReq; i++ {
		go func() {
			for url := range urls {
				timeout := time.Duration(60 * time.Second)
				client := http.Client{
					Timeout: timeout,
				}
				resp, err := client.Get(url)
				if err != nil {
					fmt.Print(err)
					continue
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Print(err)
					continue
				}
				resp.Body.Close()
				text := string(body)
				s.parse(text)
			}
			wg.Done()
		}()
	}
	for _, url := range s.StartUrls {
		urls <- url
	}
	close(urls)
	wg.Wait()
	close(s.Pipe)
}
