package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func postData() string {
	req, err := http.NewRequest("GET", "http://localhost:8999", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// set Authorization and send request
	req.Header.Set("Authorization", "secret")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	byteStr, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(byteStr)
}

func main() {
	msg := postData()
	fmt.Println(msg)
}
