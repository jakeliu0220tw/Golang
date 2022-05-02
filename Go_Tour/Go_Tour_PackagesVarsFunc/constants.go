package main

import "fmt"

const Pi = 3.14

func constants() {
	const Truth = true
	// Pi = 6.99 ==> wrong! we could not assign value to a constant
	fmt.Println("=====================")
	fmt.Println(Pi, Truth)
}
