package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	v := make(map[string]interface{})
	v["checknum"] = 123
	v["amount"] = 200
	v["category"] = []string{"gift", "clothing"}

	byteStr, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(byteStr))
}
