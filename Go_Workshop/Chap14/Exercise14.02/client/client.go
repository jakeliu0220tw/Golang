package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type response struct {
	Message string `json:"message"`
}

func getData(url string) response {
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	// must close conn at end
	defer r.Body.Close()

	// read byteStr from body
	resp := response{}
	byteStr, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	// decode to response obj
	err = json.Unmarshal(byteStr, &resp)
	if err != nil {
		fmt.Println(err)
	}

	return resp
}

func main() {
	resp := getData("http://localhost:8080")
	fmt.Println(resp.Message)
}
