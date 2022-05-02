package main

import "fmt"

type Vertex struct {
	x int
	y int
}

// struct is a collection of fields
func structFunc() {
	fmt.Println(Vertex{1, 2})
}

func structFunc2() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.x = 55
	v.y = 66
	fmt.Println(v)
	fmt.Printf("v.x = %d, v.y = %d \n", v.x, v.y)
}
