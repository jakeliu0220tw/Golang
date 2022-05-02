package main

import "fmt"

func sum(i ...int) int {
	total := 0
	for _, v := range i {
		total = total + v
	}
	return total
}

func main() {
	i := []int{100, 200, 300}
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(i...))
}
