package main

import (
	"fmt"
)

// method of type Vertex
// value receiver operates on a copy of Vertex
func (v Vertex) ScaleMethod(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

// method of type Vertex
// pointer receiver could modify the struct
func (v *Vertex) ScaleMethod2(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func methodsPointer() {
	v := Vertex{3, 4}
	v.ScaleMethod(10)
	fmt.Println(v.Abs())
	v.ScaleMethod2(10)
	fmt.Println(v.Abs())
	ScaleFunc(&v, 10)
	fmt.Println(v.Abs())
}
