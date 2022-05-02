package main

import "fmt"

func nums(i ...int) {
	fmt.Println(i)
}

func main() {
	nums(100, 200)
	nums(5)
	nums()

	s := []int{1, 2, 3, 4, 5}
	nums(s...)
}
