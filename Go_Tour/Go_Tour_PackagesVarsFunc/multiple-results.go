package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func multipleResults() {
	fmt.Println("=====================")
	var a = "hello"
	var b = "world"
	fmt.Println(a, b)
	a, b = swap(a, b)
	fmt.Println(a, b)
}
