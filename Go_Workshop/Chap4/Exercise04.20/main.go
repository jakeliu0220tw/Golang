package main

import "fmt"

func main() {
	var n name = "Golang"
	n.printName()

	p1, p2 := point{}, point{}
	p1.setPoint(10, 10)
	p2.setPoint(20, 20)
	fmt.Println("p1:", p1.getPoint())
	fmt.Println("p2:", p2.getPoint())
}
