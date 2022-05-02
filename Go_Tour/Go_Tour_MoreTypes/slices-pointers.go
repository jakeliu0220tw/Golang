package main

import "fmt"

func slicesPointers() {
	names := [4]string{
		"John",
		"Paul",
		"Goerge",
		"Ringo",
	}
	fmt.Println(names)

	// slices that share the same underlying array will see those changes
	a := names[0:2] // names[0] + names[1]
	b := names[1:3] // names[1] + names[2]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("names =", names)
}
