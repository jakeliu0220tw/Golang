package main

import "Exercise08.01/shape"

func main() {
	t := shape.Triangle{
		Base:   10,
		Height: 10,
	}
	r := shape.Rectangle{
		Length: 10,
		Width:  10,
	}
	s := shape.Square{
		Side: 10,
	}

	shape.PrintShapesDetail(t, r, s)
}
