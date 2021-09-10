package main

import "fmt"

func Sqrt(x float64) float64 {
	// z**2 should be as close to x as possible
	z := float64(x)

	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}

	return z
}

func exerciseLoopFunc() {
	fmt.Println(Sqrt(2))
}
