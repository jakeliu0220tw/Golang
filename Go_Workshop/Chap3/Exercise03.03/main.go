package main

import "fmt"

func main() {
	// demo of wraparound

	var a uint8 = 253
	var b int8 = 125
	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Println(i, ")", "uint8", a, "int8", b)
	}
}
