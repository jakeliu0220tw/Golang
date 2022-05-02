package main

import "fmt"

func incrementor() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func decrementor(max int) func() int {
	return func() int {
		max--
		if max < 0 {
			return 0
		}
		return max
	}
}

func main() {
	counterInc := incrementor()
	fmt.Println(counterInc())
	fmt.Println(counterInc())
	fmt.Println(counterInc())
	fmt.Println(counterInc())

	counterDec := decrementor(80)
	fmt.Println(counterDec())
	fmt.Println(counterDec())
	fmt.Println(counterDec())
	fmt.Println(counterDec())
}
