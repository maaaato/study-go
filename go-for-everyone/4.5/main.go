package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("defered")
	}()
	os.Exit(0)
}
