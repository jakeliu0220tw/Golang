package main

import "fmt"

func slicesFunc() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes =", primes)

	var s []int = primes[1:4] // declare a slices "s", which includes primes[1] & primes[2] & primes[3]
	fmt.Println("s =", s)
}
