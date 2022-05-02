package main

import "fmt"

type cat struct {
	name string
}

func main() {
	i := 99
	b := false
	str := "test"
	c := cat{name: "juice"}
	printDetails(i, b, str, c)
}

func printDetails(data ...interface{}) {
	for _, v := range data {
		fmt.Printf("Type:%T => Value:%v\n", v, v)
	}
}
