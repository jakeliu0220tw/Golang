package main

import "fmt"

func deferFunc() {
	// defer statement, the function call is not executed until the surrounding function returns
	defer fmt.Println("world!")

	fmt.Println("hello")
}

func deferStack() {
	fmt.Println("counting")

	// deferred call are executed in last-in-first-out order
	for i := 0; i < 10; i++ {
		defer fmt.Println("i =", i)
	}

	fmt.Println("done")
}
