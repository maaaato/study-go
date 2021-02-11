package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write(([]byte)("Hello"))

	f, err = os.Create("test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write([]byte("world"))
}
