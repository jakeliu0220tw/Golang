package shape

type Shape interface {
	area() float64
	name() string
}

func GetArea(s Shape) float64 {
	return s.area()
}

func GetName(s Shape) string {
	return s.name()
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) area() float64 {
	return t.Base * t.Height * 0.5
}

func (t Triangle) name() string {
	return "Triangle"
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) name() string {
	return "Rectangle"
}

type Square struct {
	Side float64
}

func (s Square) area() float64 {
	return s.Side * s.Side
}

func (s Square) name() string {
	return "Square"
}
