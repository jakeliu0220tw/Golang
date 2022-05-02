package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type book struct {
	ISBN          string `json:"isbn"`
	Title         string `json:"title"`
	YearPublished int    `json:"yearpublished"`
	Author        string `json:"author,omitempty"`
	CoAuthor      string `json:"coauthor,omitempty"`
}

func main() {
	b := book{}
	b.ISBN = "9910AAGG"
	b.Title = "Intro to Golang"
	b.YearPublished = 2020

	byteStr, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(byteStr))
}
