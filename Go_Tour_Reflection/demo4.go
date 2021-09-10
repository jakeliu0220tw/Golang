package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// print() prints the method set of the value x.
func print(x interface{}) {
	rv := reflect.ValueOf(x)
	rt := rv.Type()
	fmt.Printf("type %s\n", rt)

	for i := 0; i < rv.NumMethod(); i++ {
		methodType := rv.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", rt, rt.Method(i).Name,
			strings.TrimPrefix(methodType.String(), "func"))
	}
}

func demo4() {
	print(time.Hour)
	print(new(strings.Replacer))
}
