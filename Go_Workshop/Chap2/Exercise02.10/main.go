package main

import "fmt"

// func main() {
// 	config := map[string]string{
// 		"debug": "1",
// 		"info":  "2",
// 		"warn":  "3",
// 		"error": "4",
// 	}

// 	for key, value := range config {
// 		fmt.Println(key, "=>", value)
// 	}
// }

func main() {
	names := []string{"jake", "eva", "dezember"}
	for _, value := range names {
		fmt.Println(value)
	}
}
