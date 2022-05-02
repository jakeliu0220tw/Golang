package main

import "fmt"

type MyInterface interface {
	MyMethod()
}

type MyStruct struct {
	str string
}

// type "MyStruct" implement interface "MyInterface"
func (s *MyStruct) MyMethod() {
	if s == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println(s.str)
}

func helper(i MyInterface) {
	fmt.Printf("(value, type) => (%v, %T)\n", i, i)
}

func interfaceValueWithNil() {
	var i MyInterface
	var t *MyStruct

	// runtime error!
	// i.MyMethod()

	// interface value can transfer type dynamically ...
	// now i's concrete type is pointer to MyStruct type
	i = t
	helper(i)
	i.MyMethod()

	// now i's concrete type is pointer to MyStruct type
	i = &MyStruct{"HelloWorld"}
	helper(i)
	i.MyMethod()
}
