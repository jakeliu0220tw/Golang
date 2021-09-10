package main

import "fmt"

func arrayFunc() {
	var ary [2]string // declare an array "ary" with 2 string elements
	fmt.Println("ary =", ary)
	ary[0] = "Hello"
	ary[1] = "World"
	fmt.Println(ary[0], ary[1], ary)

	primes := [6]int{1, 2, 3, 4, 5, 6} // declares an array "primes" with 6 int elements
	fmt.Println("primes =", primes)
}
