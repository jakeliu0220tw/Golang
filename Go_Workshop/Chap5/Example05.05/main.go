package main

import "fmt"

func personAge(name string, age int) {
	fmt.Println(name, "is", age, "years old.")
}

func main() {
	name := "Jake"
	age := 42

	defer personAge(name, age)

	age *= 2
	personAge(name, age)
}
