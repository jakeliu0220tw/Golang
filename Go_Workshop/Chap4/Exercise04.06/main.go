package main

import "fmt"

func fillArray(arr [5]int) [5]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	return arr
}

func opArray(arr [5]int) [5]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] + 1
	}
	return arr
}

func main() {
	var arr [5]int
	arr = fillArray(arr)
	fmt.Println(arr)
}
