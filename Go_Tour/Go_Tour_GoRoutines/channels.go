package main

// send v to channel ch
// ch <- v
// receive from channel ch, and assign value to v.
// v := <-ch

// like maps, channels must be created before use
// ch := make(chan int)

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	// send sum to channel c
	c <- sum
}

func channelsFunc() {
	s := []int{7, 2, 8, -9, 4, 0}

	// create a channel c
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// receive from channel c
	// x := <-c
	// y := <-c
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
