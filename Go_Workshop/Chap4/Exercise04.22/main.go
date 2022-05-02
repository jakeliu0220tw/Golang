package main

import (
	"errors"
	"fmt"
)

func doubler(i interface{}) (string, error) {
	if v, ok := i.(int); ok {
		return fmt.Sprint(v * 2), nil
	}
	if v, ok := i.(float64); ok {
		return fmt.Sprint(v * 2), nil
	}
	if v, ok := i.(string); ok {
		return (v + v), nil
	}

	return "", errors.New("unsupported type!")
}

// demo of type assertion

func main() {
	if res, err := doubler(5); err == nil {
		fmt.Println("5 =>", res)
	}

	if res, err := doubler(3.14); err == nil {
		fmt.Println("3.14 =>", res)
	}

	if res, err := doubler("yaml"); err == nil {
		fmt.Println("yaml =>", res)
	}

	res, err := doubler(true)
	fmt.Println("true =>", res)
	fmt.Println("err =>", err)
}
