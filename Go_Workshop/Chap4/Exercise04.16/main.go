package main

import "fmt"

type id string

func getIds() (id, id, id) {
	var id1 id
	var id2 id = "123-456"
	var id3 id
	id3 = "123-456"
	return id1, id2, id3
}

func main() {
	id1, id2, id3 := getIds()
	fmt.Println("id1 == id2:", id1 == id2)
	fmt.Println("id2 == id3:", id2 == id3)
	fmt.Println("id2 == \"123-456\":", string(id2) == "123-456")
}
