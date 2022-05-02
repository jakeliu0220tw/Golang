package main

import "fmt"

func main() {
	names := []string{"jake", "eva", "dezember"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
	for idx, value := range names {
		fmt.Println(idx, "=>", value)
	}
}
