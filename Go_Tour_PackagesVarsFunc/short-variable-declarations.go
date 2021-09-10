package main

import "fmt"

func shortVariableDeclarations() {
	var i, j int = 1, 2
	k := 3
	l, m, n := true, false, "no!"
	fmt.Println("=====================")
	fmt.Println(i, j, k, l, m, n)
}
