package main

import "fmt"

func main() {
	names := "王小明"

	// print bytes length of string
	fmt.Println(len(names))

	// print length of utf-8 string
	runes := []rune(names)
	fmt.Println(len(runes))
}
