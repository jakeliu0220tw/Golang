package main

import "fmt"

// "naked" return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// return x, y
	return
}

func nakedReturn() {
	fmt.Println("=====================")
	fmt.Println(split(17))
}
