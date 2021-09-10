package main

import "fmt"

//func add(x int, y int) int {
func add(x, y int) int {
	return (x + y)
}

func functions() {
	fmt.Println("=====================")
	fmt.Println(add(42, 13))
}
