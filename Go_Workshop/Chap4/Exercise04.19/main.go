package main

import "fmt"

type name string

type location struct {
	x int
	y int
}

type size struct {
	width  int
	height int
}

type dot struct {
	name
	location
	size
}

func getDots() []dot {
	dots := []dot{}

	var dot1 dot
	dot2 := dot{}
	dot2.name = "dot2"
	dot2.x = 10
	dot2.y = 10
	dot2.width = 100
	dot2.height = 100

	// prefer this one
	dot3 := dot{
		name:     "dot3",
		location: location{x: 1, y: 1},
		size:     size{width: 1, height: 1},
	}

	// prefer this one
	dot4 := dot{}
	dot4.name = "dot4"
	dot4.location.x = -100
	dot4.location.y = -100
	dot4.size.width = 5
	dot4.size.height = 5

	dots = append(dots, dot1, dot2, dot3, dot4)
	return dots
}

func main() {
	dots := getDots()
	for i := 0; i < len(dots); i++ {
		fmt.Printf("dot:%v => %#v\n", i, dots[i])
	}
}
