package main

import "fmt"

// compare arrays
func compArrays(arr1 [5]int, arr2 [5]int, arr3 [5]int) (bool, bool, bool) {

	return arr1 == arr2, arr2 == arr3, arr3 == arr1
}

func main() {
	var arr1 [5]int
	arr2 := [5]int{0}
	arr3 := [...]int{0, 0, 0, 0, 0}

	ret1, ret2, ret3 := compArrays(arr1, arr2, arr3)
	fmt.Println(ret1)
	fmt.Println(ret2)
	fmt.Println(ret3)
}
