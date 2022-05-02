package main

import (
	"fmt"
	"time"
)

func pushData(from int, to int, ch chan int) {
	for i := from; i <= to; i++ {
		ch <- i
		time.Sleep(time.Millisecond)
	}
}

func main() {
	sum := 0
	ch := make(chan int, 100)
	defer close(ch)

	go pushData(1, 25, ch)
	go pushData(26, 50, ch)
	go pushData(51, 75, ch)
	go pushData(76, 100, ch)

	for i := 1; i <= 100; i++ {
		value := <-ch
		fmt.Println(value)
		sum += value
	}

	fmt.Println(sum)
}
