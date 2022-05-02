package main

import (
	"fmt"
	"math"
)

type F float64

// method M() of type F, means type F implement interface I
func (f F) M() {
	fmt.Println(f)
}

func interfaceValues() {
	// an interface value holds a value of a specific concrete type
	var i I

	i = &T{"Hello"}
	describeInterface(i)
	i.M()

	i = F(math.Pi)
	describeInterface(i)
	i.M()
}

func describeInterface(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
