package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("defer in main()")
	}()

	test()
	fmt.Println("don't print out")
}

func test() {
	defer func() {
		fmt.Println("defer in test()")
	}()
	message("error")
}

func message(msg string) {
	defer func() {
		fmt.Println("defer in message()")
	}()

	if msg == "error" {
		panic(errors.New("something wrong"))
	}

	fmt.Print("don't print out")
}
