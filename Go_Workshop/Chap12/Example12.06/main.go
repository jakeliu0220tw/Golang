package main

import (
	"fmt"
	"os"
)

func main() {
	txt, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(txt))
}
