package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	Id       int     `json:"id"`
	Lname    string  `json:"lname"`
	Fname    string  `json:"fname"`
	Minitial string  `json:"minitial"`
	Enrolled bool    `json:"enrolled"`
	Classes  []class `json:"classes"`
}

type class struct {
	Coursename  string `json:"coursename"`
	Coursenum   int    `json:"coursenum"`
	Coursehours int    `json:"coursehours"`
}

func main() {
	data := []byte(`
		{
			"id":123,
			"lname":"Smith",
			"minitial":null,
			"fname":"John",
			"enrolled":true,
			"classes":[
				{
					"coursename":"Intro to Golang",
					"coursenum":101,
					"coursehours":4
				},
				{
					"coursename":"English List",
					"coursenum":101,
					"coursehours":3
				},
				{
					"coursename":"World History",
					"coursenum":101,
					"coursehours":3
				}
			]
		}
	`)

	if !json.Valid(data) {
		fmt.Println("invalid json data")
		os.Exit(1)
	}

	s := student{}
	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%#v\n", s)
}
