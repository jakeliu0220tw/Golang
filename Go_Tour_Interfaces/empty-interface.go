package main

import "fmt"

func emptyInterface() {
	// i is empty interface, which handles unknown type
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

// i is empty interface. handles value of unknown type
func describe(i interface{}) {
	fmt.Printf("(value, type) => (%v, %T)\n", i, i)
}
