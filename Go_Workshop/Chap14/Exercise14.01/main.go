package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getData(url string) string {
	// get data with HttpGet
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// close connection
	defer resp.Body.Close()

	byteStr, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(byteStr)
}

func main() {
	data := getData("http://www.google.com")
	fmt.Println(data)
}
