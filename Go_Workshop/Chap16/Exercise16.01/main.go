package main

import (
	"fmt"
	"time"
)

func sum(begin int, end int) int {
	sum := 0
	for i := begin; i <= end; i++ {
		sum += begin
	}
	return sum
}

func main() {
	var s1, s2 int
	go func() {
		s1 = sum(1, 10)
	}()

	s2 = sum(1, 100)
	time.Sleep(time.Second)

	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
}
