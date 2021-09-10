package main

import "fmt"

// var m map[string]Vertex

func mapsFunc() {
	// initialize a map "m"
	m := make(map[string]Vertex)
	m["bell labs"] = Vertex{x: 40, y: -74}
	fmt.Println("m[\"bell labs\"] =", m["bell labs"])
}
