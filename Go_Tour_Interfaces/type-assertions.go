package main

import "fmt"

func typeAssertions() {
	// i is empty interface value, which handles unknow type
	var i interface{} = "hello"

	// assign string value from interface value, i to s
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// assign float64 value from interface value i to f
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// panic !!
	// f = i.(float64)
}
