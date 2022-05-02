package main

import "fmt"

func mapLiterals() {
	m := make(map[string]Vertex)
	m["BellLabs"] = Vertex{40, -74}
	m["Google"] = Vertex{37, -122}

	fmt.Println(m)
}
