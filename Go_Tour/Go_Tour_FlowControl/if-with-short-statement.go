package main

import (
	"fmt"
	"math"
)

func pow(x, y, lim float64) float64 {
	// v := math.Pow(x, y)
	// if v < lim {
	// 	return v
	// }

	if v := math.Pow(x, y); v < lim {
		return v
	}

	return lim
}

func pow2(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		fmt.Printf("%g < %g, return %g\n", v, lim, v)
		return v
	} else {
		fmt.Printf("%g >= %g, return %g\n", v, lim, lim)
		return lim
	}
}

func ifWithShortStatement() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}
