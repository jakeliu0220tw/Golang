package main

import "fmt"

// declare an interface, I, which should implement M() method
type I interface {
	M()
}

// declare a struct, T, which contains string S
type T struct {
	S string
}

// method M() of type T, means type T implements the interface I.
func (t *T) M() {
	fmt.Println(t.S)
}

func interfacesImplementation() {
	var myInterface I = &T{"Hello World"}
	myInterface.M()
}
