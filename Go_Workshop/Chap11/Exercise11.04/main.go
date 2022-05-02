package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type student struct {
	StudentId  int    `json:"studentid"`
	LastName   string `json:"lname"`
	FirstName  string `json:"fname"`
	IsEnrolled bool   `json:"isenrolled"`
	Age        int    `json:"age"`
}

func main() {
	s := student{
		StudentId:  123,
		LastName:   "Liu",
		FirstName:  "Jake",
		IsEnrolled: true,
		Age:        42,
	}
	s2 := student{}

	// encode student obj to bytes buffer
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(&s); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%x\n", buffer)

	// decode bytes buffer to student obj
	decoder := gob.NewDecoder(&buffer)
	if err := decoder.Decode(&s2); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(s2)
}
