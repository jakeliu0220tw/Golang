package main

import "fmt"

type Vortex struct {
	x int
	y int
}

func structPointers() {
	v := Vortex{1, 2}
	fmt.Println(v)
	p := &v     // pointer to struct v
	(*p).x = 11 // formal form
	p.y = 22    // short form
	fmt.Println(*p)
}
