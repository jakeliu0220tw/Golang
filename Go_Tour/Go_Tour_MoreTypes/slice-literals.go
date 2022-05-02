package main

import "fmt"

func sliceLiterals() {
	// declare a slice q, which automatically allocate an array with 6 int elements
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("q =", q)

	// declare a slice r, which automatically allocate an array with 6 bool elements
	r := []bool{true, false, true, true, false, true}
	fmt.Println("r =", r)

	type myStruct struct {
		i int
		b bool
	}

	// declare a slice "s", which allocate an array with 6 myStruct elements
	s := []myStruct{
		{i: 2, b: true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println("s =", s)
}
