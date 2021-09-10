package main

import "fmt"

// non-struct type, MyFloat
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func methodsContinued() {
	f := MyFloat(-12)
	ff := MyFloat(12)
	fmt.Println(f.Abs())
	fmt.Println(ff.Abs())
}
