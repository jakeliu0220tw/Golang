package main

// https://pkg.go.dev/gopkg.in/yaml.v2

import (
	"fmt"
	"io/ioutil"
	"log"
)

func WriteFile() {
	// message := []byte("Hello, Gophers!")
	// err := ioutil.WriteFile("hello", message, 0644)
	fmt.Println("write data into file")
	err := ioutil.WriteFile("hello", []byte("hello, world"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile() {
	content, err := ioutil.ReadFile("hello")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s\n", content)
}

func main() {
	fmt.Println("====================")
	WriteFile()
	fmt.Println("====================")
	ReadFile()
}
