package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	defer fmt.Println("done")
	trapSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, trapSignals...)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sig := <-sigCh
		fmt.Println("Got signal", sig)
		cancel()
	}()
	doMain(ctx)
}

func doMain(ctx context.Context) {
	defer fmt.Println("done doMain")
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
