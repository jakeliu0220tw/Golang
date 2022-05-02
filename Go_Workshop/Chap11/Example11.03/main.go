package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Lastname  string  `json:"lname"`
	Firstname string  `json:"fname"`
	Address   address `json:"address"`
}

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode int    `json:"zipcode"`
}

func main() {
	data := []byte(`
		{
			"lname":"smith",
			"fname":"john",
			"address":
			{
				"street":"xxx rd",
				"city":"park city",
				"state":"CA",
				"zipcode":123456
			}
		}
	`)

	if !json.Valid(data) {
		fmt.Println("invalid json format")
		os.Exit(1)
	}

	v := person{}
	if err := json.Unmarshal(data, &v); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", v)
	}
}
