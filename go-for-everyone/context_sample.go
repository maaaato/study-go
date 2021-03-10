package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go fetchURL(ctx, queue)
	}

	// 送信
	queue <- "https://www.example.com"
	queue <- "https://www.example.net"
	queue <- "https://www.example.net/foo"
	queue <- "https://www.exmaple.net/bar"

	cancel()
	wg.Wait()
}

func fetchURL(ctx context.Context, queue chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker exit")
			wg.Done()
			return
		case url := <-queue:
			fmt.Println("fetching", url)
		}
	}
}
