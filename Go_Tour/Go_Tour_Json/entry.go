package main

import (
	"encoding/json"
	"fmt"
)

// decode a json blob from internal objs
func unmarshalDemo() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)

	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal

	// func Unmarshal(data []byte, v interface{}) error
	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}

	// output animals[]
	for _, v := range animals {
		fmt.Printf("%+v\n", v)
	}
}

// encode a internal obj into a json blob
func marshalDemo() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))
}

func main() {
	fmt.Println("====================")
	unmarshalDemo()
	fmt.Println("====================")
	marshalDemo()
}
