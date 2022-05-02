package main

import "fmt"

var budgetMap = make(map[int]string)

func init() {
	fmt.Println("init")
	budgetMap[1] = "food"
	budgetMap[2] = "clothes"
	budgetMap[3] = "house"
	budgetMap[4] = "traffic"
}

func main() {
	fmt.Println("main")
	for k, v := range budgetMap {
		fmt.Printf("%v => %v\n", k, v)
	}
}
