package main

import "fmt"

func mutatingMaps() {
	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println("The value:", m["answer"])

	m["answer"] = 48
	fmt.Println("The value:", m["answer"])

	// delete an element from map
	delete(m, "answer")
	fmt.Println("The value:", m["answer"])

	// test if the key is present
	val, exist := m["answer"]
	fmt.Println("The value:", val, "Present?", exist)
}
