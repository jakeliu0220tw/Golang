package main

import (
	"encoding/json"
	"fmt"
)

type greeting struct {
	// field name must be Capital otherwise it will not be exported by Unmarshal()
	Message string
}

func main() {
	var v greeting
	data := []byte(`{"Message":"greeting gopher!"}`)

	err := json.Unmarshal(data, &v)
	if err == nil {
		fmt.Println(data)
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}
}
