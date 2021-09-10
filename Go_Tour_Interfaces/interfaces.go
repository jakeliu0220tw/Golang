package main

import (
	"fmt"
	"math"
)

// interface type is defined as a set of method signatures
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func interfacesFunc() {
	var a Abser
	f := MyFloat(-21)
	v := Vertex{3, 4}

	// a MyFloat implements Abs()
	a = f
	fmt.Println(a.Abs())
	// a *Vertex implements Abs()
	a = &v
	fmt.Println(a.Abs())

	// compile error! because Vertex does not implement Abs()
	// a = v
}
