package main

import "fmt"

// sender
func fibonacci(n int, c chan int) {
	fmt.Println("This is a sender")
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Note: only the sender should close a channel
	// close channel c
	close(c)
}

func rangeAndClose() {
	// create a channel c
	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	// the loop "for i := range c" receives values from the channel repeatedly until it is closed
	// receiver
	fmt.Println("This is a receiver")
	for i := range c {
		fmt.Println(i)
	}
}
