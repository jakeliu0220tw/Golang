package main

import (
	"fmt"
	"math"
)

// arr1 = [0, 3, 5, 2, 9, -10]
// arr2 = [4, 1, 2, 10, 5, 20]
// target: 24
// ans: [5, 20]
// given 2 integer array
// find i, j such that minimize abs(arr1[i]+arr2[j]-target)

// method 1: brute force time->O(n^2), space O(1)
// any other method ???

func main() {
	arr1 := []int{0, 3, 5, 2, 9, -10}
	arr2 := []int{4, 1, 2, 10, 5, 20}

	ans1 := closestSum(arr1, arr2, 24)
	fmt.Println(ans1)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func closestSum(arr1, arr2 []int, target int) []int {
	if len(arr1) == 0 && len(arr2) == 0 {
		return []int{}
	}

	if len(arr1) == 1 && len(arr2) == 1 {
		return []int{arr1[0], arr2[0]}
	}

	diff := math.MaxInt32
	arr1Idx := 0
	arr2Idx := 0
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if Abs(arr1[i]+arr2[j]-target) <= diff {
				diff = Abs(arr1[i] + arr2[j] - target)
				arr1Idx = i
				arr2Idx = j
			}
		}
	}

	return []int{arr1[arr1Idx], arr2[arr2Idx]}
}
