package main

import "fmt"

func forFunc() {
	sum := 0
	// for loop in go, unlike C & Java, no need parentheses
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println("sum = ", sum)

	// while loop in go, loop will continue if sum2 < 1000, it's similar with "while" condition in C / C++
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println("sum2 = ", sum2)
}
