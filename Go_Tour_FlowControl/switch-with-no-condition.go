package main

import (
	"fmt"
	"time"
)

// this is a good form to replace long if-elseif-else statement.
func switchWithNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning!")
	case t.Hour() < 17:
		fmt.Println("good afternoon!")
	default:
		fmt.Println("good night!")
	}
}
