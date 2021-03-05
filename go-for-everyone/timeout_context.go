package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func getHTTP(url string, dst io.Writer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(dst, resp.Body)
	return err
}

func main() {
	dst, _ := os.Create("context.txt")
	err := getHTTP("https://yahoo.com", dst)
	if err != nil {
		fmt.Println("aaa")
	}
}
