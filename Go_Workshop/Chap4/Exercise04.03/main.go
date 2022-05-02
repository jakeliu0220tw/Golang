package main

import "fmt"

func main() {
	// init array with values
	var arr1 [5]int
	fmt.Println(arr1)

	arr2 := [5]int{0, 1, 2, 3, 4}
	fmt.Println(arr2)

	arr3 := [5]int{4: 99}
	fmt.Println(arr3)

	arr4 := [5]int{1: 1, 4: 4}
	fmt.Println(arr4)

	arr5 := [...]int{11, 22, 33, 44, 55}
	fmt.Println(arr5)
}
