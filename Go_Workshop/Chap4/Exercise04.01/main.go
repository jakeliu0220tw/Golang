package main

import "fmt"

func defineArray() [10]int {
	// var arr [10]int
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return arr
}

func main() {
	fmt.Printf("%#v", defineArray())
}
