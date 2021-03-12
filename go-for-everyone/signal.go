package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type MySignal struct {
	message string
}

func (s MySignal) String() string {
	return s.message
}

func (s MySignal) Signal() {}

func main() {
	log.Println("[info] Start")
	// defer fmt.Println("done")
	trapSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT}
	// 受信するチャンネルを用意
	sigCh := make(chan os.Signal, 1)

	// 10秒後にsigChにMySignalの値を送信
	time.AfterFunc(20*time.Second, func() {
		sigCh <- MySignal{"timed out"}
	})

	signal.Notify(sigCh, trapSignals...)

	// 受信するまで待ち受ける
	sig := <-sigCh
	switch s := sig.(type) {
	case syscall.Signal:
		// osからのシグナルの場合
		log.Printf("[info] Got Signal:%s(%d)", s, s)
	case MySignal:
		// アプリケーション独自のシグナルの場合
		log.Printf("[info] %$", s) // .String()が評価される
	}
	// ctx, cancel := context.WithCancel(context.Background())
	// go func() {
	// 	sig := <-sigCh
	// 	fmt.Println("Got signal", sig)
	// 	cancel()
	// }()
	// doMain(ctx)
}

// func doMain(ctx context.Context) {
// 	defer fmt.Println("done doMain")
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 		}
// 	}
// }
