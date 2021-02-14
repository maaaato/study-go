package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Fprintln(os.Stdout, strings.Repeat("x", 100))
	}
}
