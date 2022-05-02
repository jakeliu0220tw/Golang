package main

import "fmt"

func main() {
	input := 20

	if input < 0 {
		fmt.Println("input could not be negative.")
	} else if input%2 == 0 {
		fmt.Println("input is even.")
	} else {
		fmt.Println("input is odd.")
	}
}
