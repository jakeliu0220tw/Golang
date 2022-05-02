package main

import (
	"errors"
	"fmt"
)

func doubler(i interface{}) (string, error) {
	switch t := i.(type) {
	case bool:
		return "truefalse", nil
	case string:
		return (t + t), nil
	case int8:
		return fmt.Sprint(t * 2), nil
	case int16:
		return fmt.Sprint(t * 2), nil
	case int32:
		return fmt.Sprint(t * 2), nil
	case int64:
		return fmt.Sprint(t * 2), nil
	case float32:
		return fmt.Sprint(t * 2), nil
	case float64:
		return fmt.Sprint(t * 2), nil
	default:
		return "", errors.New("unsupported type!")
	}
}

func main() {
	res, err := doubler(-5)
	fmt.Println("-5 =>", res, err)
	res, err = doubler(5)
	fmt.Println("5 =>", res, err)
	res, err = doubler("mur")
	fmt.Println("mur =>", res, err)
	res, err = doubler(true)
	fmt.Println("true =>", res, err)
	res, err = doubler(float64(3.14))
	fmt.Println("3.14 =>", res, err)
}
