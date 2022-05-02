package main

import (
	"fmt"
	"math"
)

func typeConversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println("=====================")
	fmt.Println(x, y, z)
}
