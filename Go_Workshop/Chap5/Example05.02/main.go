// package main

// import "fmt"

// type calc func(int, int) int

// func calculator(f calc, i int, j int) {
// 	ret := f(i, j)
// 	fmt.Println(ret)
// }

// func add(x, y int) int {
// 	fmt.Println(x, "add", y)
// 	return (x + y)
// }

// func subtract(x, y int) int {
// 	fmt.Println(x, "subtract", y)
// 	return (x - y)
// }

// func main() {
// 	calculator(add, 3, 5)
// 	calculator(subtract, 10, 5)
// }

package main

import "fmt"

type calc func(int, ...int) int

func calculator(f calc, i int, j ...int) {
	ret := f(i, j...)
	fmt.Println(ret)
}

func add(x int, y ...int) int {
	fmt.Println(x, "add", y)
	total := x
	for _, v := range y {
		total = total + v
	}
	return total
}

func subtract(x int, y ...int) int {
	fmt.Println(x, "subtract", y)
	total := x
	for _, v := range y {
		total = total - v
	}

	return total
}

func main() {
	calculator(add, 1, 3, 5)
	calculator(subtract, 10, 5, 3)
}
