package main

import "fmt"

var level string = "pkg"

func main() {
	fmt.Println("main start:", level)

	level := 42
	if true {
		fmt.Println("block start:", level)
	}

	myFunc()

	fmt.Println("main end:", level)
}

func myFunc() {
	fmt.Println("myFunc start:", level)
}
