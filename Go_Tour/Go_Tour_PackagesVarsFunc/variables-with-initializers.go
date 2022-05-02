package main

import "fmt"

var i, j int = 1, 2

func variablesWithInitializers() {
	var c, python, java = true, false, "NoNo"
	fmt.Println("=====================")
	fmt.Println(i, j, c, python, java)
}
