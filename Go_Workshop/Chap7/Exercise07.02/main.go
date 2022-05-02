package main

import (
	"fmt"
	"math"
)

// Shap interface
type Shap interface {
	Name() string
	Area() float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type triangle struct {
	base   float64
	height float64
}

func printShapeDetails(shapes ...Shap) {
	for _, shap := range shapes {
		fmt.Printf("name:%v => area:%v\n", shap.Name(), shap.Area())
	}
}

func main() {
	c := circle{radius: 2}
	s := square{side: 2}
	t := triangle{base: 5, height: 5}

	printShapeDetails(c, s, t)
}

// circle implement Shape interface
func (c circle) Name() string {
	return "circle"
}
func (c circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

// square implement Shape interface
func (s square) Name() string {
	return "square"
}

func (s square) Area() float64 {
	return s.side * s.side
}

// triangle implement Shape interface
func (t triangle) Name() string {
	return "triangle"
}

func (t triangle) Area() float64 {
	return t.base * t.height * 0.5
}
