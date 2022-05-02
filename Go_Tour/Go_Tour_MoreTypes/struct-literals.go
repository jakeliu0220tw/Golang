package main

import "fmt"

var (
	v1 = Vortex{1, 2}  // has type Vortex
	v2 = Vortex{x: 1}  // y:0 is implicit
	v3 = Vortex{}      // x:0 and y:0
	p  = &Vortex{3, 4} // pointer to struct Vortex
)

func structLiterals() {
	fmt.Println(v1, v2, v3, p, *p)
}
