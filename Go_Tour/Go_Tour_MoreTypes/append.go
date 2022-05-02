package main

func appendFunc() {
	// init a nil slice
	var s []int
	printSlice(s)

	// append() works on nil slice
	s = append(s, 0)
	printSlice(s)

	// the slice grows as needed
	s = append(s, 1)
	printSlice(s)

	// we can add more than one elements at a time
	s = append(s, 2, 3, 4)
	printSlice(s)
}
