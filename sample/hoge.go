package main

import (
	"fmt"
	"os"
)

func main() {
	go func() {
		fmt.Println("aaa")
		os.Exit(0)
	}()

	for {
	}
}
