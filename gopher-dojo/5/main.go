package main

import (
	"flag"
	"fmt"
	"strings"
)

var msg = flag.String("msg", "default", "説明")

var n int

func init() {
	flag.IntVar(&n, "n", 1, "回数")
}

func main() {
	flag.Parse()
	var files = flag.Arg(0)
	fmt.Println(strings.Repeat(*msg, n))
	fmt.Println(files)
}
