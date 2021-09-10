package main

import (
	"fmt"
	"reflect"
)

// we could use reflect module to retrieve tags
type T struct {
	f string `one:"1" two:"2" blank:""`
}

func demo2() {
	t := T{}
	fmt.Println("t =", t)

	tr := reflect.TypeOf(t)
	fmt.Println("tr =", tr)

	f, _ := tr.FieldByName("f")
	fmt.Println("f =", f)
	fmt.Println("f.Tag =", f.Tag)

	v, ok := f.Tag.Lookup("one")
	fmt.Printf("%s, %t\n", v, ok)

	v, ok = f.Tag.Lookup("two")
	fmt.Printf("%s, %t\n", v, ok)

	v, ok = f.Tag.Lookup("three")
	fmt.Printf("%s, %v\n", v, ok)
}
