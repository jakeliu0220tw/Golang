package main

import "fmt"

func main() {
	// raw string
	comment1 := `This is the best
	thing ever.`
	fmt.Println(comment1)

	comment2 := `This is the best thing ever.`
	fmt.Println(comment2)

	comment3 := `This is the best\n thing ever.`
	fmt.Println(comment3)

	comment4 := "This is the best\n thing ever."
	fmt.Println(comment4)
}
