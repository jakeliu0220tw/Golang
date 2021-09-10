package main

import (
	"encoding/json"
	"fmt"
)

// json marshal is one example of tags
type MyJson struct {
	F1 int `json:"f1"`
	F2 int `json:"f2"`
	F3 int `json:"f3"`
	F4 int `json:"f4"`
}

func demo3() {
	t := MyJson{0, 2, 1, 3}
	fmt.Println("t =", t)
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json string of t: %s\n", b)
}
