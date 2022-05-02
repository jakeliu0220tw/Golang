package main

import "fmt"

// define a function type
type calc func(int, int) int

func add(i, j int) int {
	return (i + j)
}

func calculator(f calc, i int, j int) {
	fmt.Println(i, "+", j, "=", f(i, j))
}

func main() {
	calculator(add, 10, 20)
}
