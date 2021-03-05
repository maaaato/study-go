package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func getHTTP(url string, dst io.Writer) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(dst, resp.Body)
	return err
}

func main() {
	dst, _ := os.Create("res.txt")
	err := getHTTP("https://google.com", dst)
	if err != nil {
		fmt.Println("aaa")
	}
}
