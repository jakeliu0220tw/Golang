package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// iterate way to show each elements of slice
func rangeFunc() {
	// two value will return for each iteration, first is index, second is a copy of the elements
	for i, v := range pow {
		fmt.Printf("2**%d = %d \n", i, v)
	}
}
