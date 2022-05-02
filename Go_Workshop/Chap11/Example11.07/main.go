package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct {
	LastName  string `json:"lname"`
	FirstName string `json:"fname"`
	Address   string `json:"address"`
}

func main() {
	byteStr := []byte(`
		{
			"lname": "liu",
			"fname": "jake",
			"address": "Far far away island"
		}
	`)
	p := person{}

	// new a json decoder
	decoder := json.NewDecoder(strings.NewReader(string(byteStr)))
	// decode the byteStr to person obj
	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", p)

	// new a json encoder
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(p); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
