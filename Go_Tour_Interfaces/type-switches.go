package main

import "fmt"

func do(i interface{}) {
	// here is the example of type switch
	switch v := i.(type) {
	case int:
		fmt.Printf("type int: %v * 2 = %v\n", v, v*2)
	case string:
		fmt.Printf("type string: %q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("type unknown: I don't know about the type ... %T\n", v)
	}
}

func typeSwitches() {
	do(21)
	do("hello")
	do(true)
}
