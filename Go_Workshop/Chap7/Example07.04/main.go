package main

import "fmt"

type cat struct {
	name string
}

func main() {
	i := []interface{}{"Jake", 42, cat{name: "juice"}, false}
	printType(i)

	printTypes("Jake", 42, cat{name: "juice"}, false)
}

func printType(types []interface{}) {
	for _, v := range types {
		switch x := v.(type) {
		case int:
			fmt.Printf("%v is int\n", x)
		case bool:
			fmt.Printf("%v is bool\n", x)
		case string:
			fmt.Printf("%v is string\n", x)
		default:
			fmt.Printf("%v is unknown type\n", x)
		}

	}
}

func printTypes(types ...interface{}) {
	for _, v := range types {
		switch x := v.(type) {
		case int:
			fmt.Printf("%v is int\n", x)
		case bool:
			fmt.Printf("%v is bool\n", x)
		case string:
			fmt.Printf("%v is string\n", x)
		default:
			fmt.Printf("%v is unknown type\n", x)
		}

	}
}
