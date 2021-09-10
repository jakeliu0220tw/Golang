package main

import "fmt"

func zero() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println("=====================")
	// %q is string
	fmt.Printf("%v, %v, %v, %q\n", i, f, b, s)
}
