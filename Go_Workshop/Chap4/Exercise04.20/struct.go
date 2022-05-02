package main

import "fmt"

type name string

// method of name struct
func (n name) printName() {
	fmt.Println("name:", n)
}

type point struct {
	x int
	y int
}

// mehtod of point struct
func (p *point) getPoint() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func (p *point) setPoint(x, y int) {
	p.x = x
	p.y = y
}
