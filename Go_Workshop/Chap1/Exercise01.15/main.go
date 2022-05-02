package main

import "fmt"

func add5ByValue(count int) {
	count += 5
	fmt.Printf("add5ByValue: %#v\n", count)
}
func add5ByPtr(count *int) {
	*count += 5
	fmt.Printf("add5ByPtr: %#v\n", *count)
}

func main() {
	var count int = 0
	fmt.Printf("before add5ByValue: %#v\n", count)
	add5ByValue(count)
	fmt.Printf("before add5ByPtr: %#v\n", count)
	add5ByPtr(&count)
}
