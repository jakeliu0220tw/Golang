package main

import "fmt"

func main() {
	add := calculator("+")
	substract := calculator("-")
	fmt.Println("3+5=", add(3, 5))
	fmt.Println("10-5=", substract(10, 5))
}

func calculator(operator string) func(int, int) int {
	switch operator {
	case "+":
		return func(x int, y int) int {
			return (x + y)
		}
	case "-":
		return func(x, y int) int {
			return (x - y)
		}
	default:
		return nil
	}
}
