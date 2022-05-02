package main

import "fmt"

func nilSlice() {
	// s is nil slice
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("s is nil slice!")
	}
}
