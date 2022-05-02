package shape

import "fmt"

type Shape interface {
	area() float64
	name() string
}

type Triangle struct {
	Base   float64
	Height float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Square struct {
	Side float64
}

func (t Triangle) area() float64 {
	return t.Base * t.Height * float64(0.5)
}

func (t Triangle) name() string {
	return "Triangle"
}

func (r Rectangle) area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) name() string {
	return "Rectangle"
}

func (s Square) area() float64 {
	return float64(s.Side * s.Side)
}

func (s Square) name() string {
	return "Square"
}

func PrintShapesDetail(shapes ...Shape) {
	for _, v := range shapes {
		fmt.Printf("%v => area:%v\n", v.name(), v.area())
	}
}
