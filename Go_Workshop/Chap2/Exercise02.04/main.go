package main

import (
	"errors"
	"fmt"
)

func validate(input int) error {
	if input < 0 {
		return errors.New("input < 0")
	} else if input > 100 {
		return errors.New("input > 100")
	} else if input%7 == 0 {
		return errors.New("input%7 == 0")
	}

	return nil
}

func main() {
	input := 26
	if err := validate(input); err != nil {
		fmt.Println("Error:", err)
	} else if input%2 == 0 {
		fmt.Println("input is even.")
	} else {
		fmt.Println("input is odd.")
	}
}
