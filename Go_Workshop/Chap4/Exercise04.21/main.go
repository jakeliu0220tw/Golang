package main

import (
	"fmt"
	"math"
)

func convert() {
	var i8 int8 = math.MaxInt8
	var i int = 128
	var f64 float64 = 3.14

	fmt.Printf("int8=%v -> int64=%v\n", i8, int64(i8))
	fmt.Printf("int64=%v -> int8=%v\n", i, int8(i))
	fmt.Printf("int8=%v -> float64=%v\n", i8, float64(i8))
	fmt.Printf("float64=%v -> int8=%v\n", f64, int8(f64))
}

func main() {
	convert()
}
