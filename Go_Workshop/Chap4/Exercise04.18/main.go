package main

import "fmt"

type point struct {
	x int
	y int
}

func comp() (bool, bool) {
	point1 := struct {
		x int
		y int
	}{
		10,
		5,
	}

	point2 := point{10, 10}

	var point3 point
	point3.x = 10
	point3.y = 10

	return point1 == point2, point2 == point3
}

func main() {
	ret1, ret2 := comp()
	fmt.Println("point1 == point2", ret1)
	fmt.Println("point2 == point3", ret2)
}
