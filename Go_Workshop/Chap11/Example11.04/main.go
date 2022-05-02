package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type greeting struct {
	Message string `json:"message"`
}

// demo of marshal()
func main() {
	s := greeting{}
	s.Message = "hello gopher!"

	json, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("json:%s\n", string(json))
}
