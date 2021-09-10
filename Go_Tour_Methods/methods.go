package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x float64
	y float64
}

// Abs() is the method of type Vertex
// a method is a function with special receiver argument
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// a normal function
func FuncAbs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func methods() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(FuncAbs(v))
}
