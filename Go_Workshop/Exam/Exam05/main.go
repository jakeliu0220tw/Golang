package main

import "fmt"

func calc(a, b int) int {
	fmt.Printf("%d+%d=%d\n", a, b, a+b)
	return (a + b)
}

func main() {
	a := 1
	b := 2

	defer calc(a, calc(a, b))
	a = 10
	defer calc(calc(a, b), b)
}

// output
// 1+2=3
// 10+2=12
// 12+2=14
// 1+3=4
// the parameters passing into defer() will be calculated first
