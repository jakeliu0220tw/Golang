package main

import "fmt"

func main() {
	query, limit, offset := "bat", 10, 0
	fmt.Println(query, limit, offset)

	query, limit, offset = "get", 200, 100
	fmt.Println(query, limit, offset)
}
