package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:",omitempty"`
	ZipCode int    `json:"-"`
}

func main() {

	addr := address{}
	addr.Street = "Galaxy Far Away"
	addr.City = "Dark City"
	addr.ZipCode = 777

	byteStr, err := json.Marshal(addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(byteStr))

	byteStr, err = json.MarshalIndent(addr, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(byteStr))

}
