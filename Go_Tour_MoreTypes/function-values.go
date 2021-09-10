package main

import (
	"fmt"
	"math"
)

// compute() receives a function ptr as parameters
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func funcValues() {
	// declare a function variable, "hypot"
	hypot := func(x float64, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
