package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type greeting struct {
	SomeMsg string `json:"message"`
}

func main() {
	v := greeting{}
	data := []byte(`{"message":"hello gopher!"}`)

	if !json.Valid(data) {
		fmt.Println("invalid json format")
		os.Exit(1)
	}

	if err := json.Unmarshal(data, &v); err != nil {
		fmt.Println(err)
	}

	fmt.Println(v)
}
