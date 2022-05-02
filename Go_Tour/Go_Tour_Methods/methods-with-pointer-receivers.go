package main

import (
	"fmt"
	"math"
)

func (v *Vertex) AbsMethod() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func methodsWithPointerReceivers() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.AbsMethod())
	v.ScaleMethod2(10)
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.AbsMethod())
}
