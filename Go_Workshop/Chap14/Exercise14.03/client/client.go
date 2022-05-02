package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type msgData struct {
	Message string `json:"message"`
}

func postData(data msgData) msgData {
	byteStr, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	buffer := bytes.NewBuffer(byteStr)

	r, err := http.Post("http://localhost:8080", "application/json", buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer r.Body.Close()

	msg := msgData{}
	err = json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return msg
}

func main() {
	msg := msgData{Message: "hello"}
	resp := postData(msg)
	fmt.Println(resp)
}
