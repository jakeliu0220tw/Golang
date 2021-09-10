package main

import (
	"fmt"
	"time"
)

// type error interface {
// 	Error() string
// }

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error happened at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it doesn't work"}
}

func errorsFunc() {
	// a nil error denotes success, a non-nil error denotes failure
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
