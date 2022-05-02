package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer#1")
	}()
	defer func() {
		fmt.Println("defer#2")
	}()
	defer func() {
		fmt.Println("defer#3")
	}()

	f1 := func() {
		fmt.Println("f1")
	}

	f2 := func() {
		fmt.Println("f2")
	}

	f1()
	f2()
}
