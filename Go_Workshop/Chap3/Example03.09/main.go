package main

import "fmt"

func main() {
	name := `劉小毛` // actually string is byte slices
	fmt.Println(name)

	for i := 0; i < len(name); i++ {
		// print byte slices
		fmt.Println(name[i], "")
	}

	for i, v := range name {
		// print utf-8 string
		fmt.Println(i, string(v))
	}
}
