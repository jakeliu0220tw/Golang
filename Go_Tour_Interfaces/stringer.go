package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// type Person implement String()
func (p Person) String() string {
	return fmt.Sprintf("%v : %v", p.Name, p.Age)
}

func stringer() {
	a := Person{"Jake", 41}
	fmt.Println(a.String())
	fmt.Println(a)

	b := Person{"Wilhelm", 41}
	fmt.Println(b)
}
