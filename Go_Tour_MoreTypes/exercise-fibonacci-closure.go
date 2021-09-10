package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	count := 0
	previous := 1
	sum := 0
	return func() int {
		count += 1
		if count == 1 {
			return 0
		}
		if count == 2 {
			return 1
		}
		if count >= 3 {
			tmp := previous
			previous = sum
			sum = tmp + sum
		}

		return (previous + sum)
	}
}

func exerciseFunctionClosure() {
	fn := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}
