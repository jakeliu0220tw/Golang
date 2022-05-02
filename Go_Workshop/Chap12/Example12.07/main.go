package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	text, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(text))
}
