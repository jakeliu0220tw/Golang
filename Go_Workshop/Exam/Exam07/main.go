package main

import "fmt"

// Given an integer array A, write a function to check if there exists two different index pair (i, j) such that A[i] = 2*A[j]
// if exist such pair, return true
// otherwise, return false

// find arr[i] == 2*arr[j]
// if yes, return true

// method 1:
// time complexity -> O(n), space complexity -> O(n)
func findAns(arr []int) bool {
	// establish myMap, key is element of arr
	myMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		_, exist := myMap[v]
		if !exist {
			myMap[v] += 1
		}
	}
	// fmt.Println(myMap)

	// scan array to find out if condition is satisfied
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		if _, exist := myMap[v*2]; exist {
			// fmt.Printf("i:%v, v:%v\n", i, v)
			return true
		}
	}

	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	ans := findAns(arr)
	fmt.Println(ans)

	arr2 := []int{-1, -2, -3, 4, 5, 6}
	ans2 := findAns(arr2)
	fmt.Println(ans2)

	arr3 := []int{-1, 2, -3}
	ans3 := findAns(arr3)
	fmt.Println(ans3)
}
