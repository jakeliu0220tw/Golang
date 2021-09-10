package main

import "fmt"

func rangeContinued() {
	pow := make([]int, 10)

	// i is index, value is omitted
	for i := range pow {
		// == 2**i
		pow[i] = 1 << uint(i)
	}

	// index is omitted, v is copy of slice element
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
