package main

import (
	"errors"
	"fmt"
)

func a() {
	b("bye")
}

func b(msg string) {
	defer func() {
		// get panic
		if r := recover(); r != nil {
			fmt.Println("get some error:", r)
		}
	}()

	if msg == "bye" {
		panic(errors.New("bye"))
	}
	fmt.Println(msg)
}

// demo of panic & recover
func main() {
	a()
	fmt.Println("main()")
}
