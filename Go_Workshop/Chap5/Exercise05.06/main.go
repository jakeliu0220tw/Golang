package main

import "fmt"

func main() {
	x := 9
	// anonymous function
	sqr := func(i int) int {
		return i * i
	}

	fmt.Println(x, "=>", sqr(x))
}
