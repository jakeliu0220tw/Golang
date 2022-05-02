package main

import "fmt"

func bufferedChannels() {
	// create a channel ch
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
