package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var n bool

func init() {
	flag.BoolVar(&n, "n", false, "Output line")
}

func main() {
	flag.Parse()
	var files = flag.Args()
	var line int
	for _, v := range files {
		f, err := os.Open(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, "can't open file:", err)
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line += 1
			if n {
				fmt.Fprintln(os.Stdout, line, ":", scanner.Text())
			} else {
				fmt.Fprintln(os.Stdout, scanner.Text())
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
