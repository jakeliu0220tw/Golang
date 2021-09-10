package main

// https://yourbasic.org/golang/iota/

import (
	"fmt"
)

// demo1() for enum usage
const (
	C0 = iota + 1
	C1
	C2
)

func demo1() {
	fmt.Println("C0 =", C0)
	fmt.Println("C1 =", C1)
	fmt.Println("C2 =", C2)
}

// demo2() for enum usage
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func demo2() {
	var d Direction = East
	fmt.Println("d is", d)

	switch d {
	case North:
		fmt.Println("goes up.")
	case South:
		fmt.Println("goes down.")
	default:
		fmt.Println("stays put.")
	}
}

func main() {
	demo1()
	demo2()
}
