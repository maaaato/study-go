package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		fmt.Println(i)
		wg.Add(1)
		go fetchURL(queue)
	}
	// 送信
	queue <- "https://www.example.com"
	queue <- "https://www.example.net"
	queue <- "https://www.example.net/foo"
	queue <- "https://www.exmaple.net/bar"
	// TODO: channelの送信のブロックについて調べる

	fmt.Println("aaa")
	close(queue)
	wg.Wait()
}

func fetchURL(queue chan string) {
	for {
		url, more := <-queue // 受信
		fmt.Println("url more", url, more)
		if more {
			fmt.Println("fetching", url)
		} else {
			fmt.Println("worker exit")
			wg.Done()
			return
		}
	}
}
