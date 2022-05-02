package main

import "fmt"

// doSomething() is some kind of consumer
func doSomething() (chan int, chan int) {
	in := make(chan int)
	out := make(chan int)

	go func() {
		sum := 0
		for i := range in {
			sum += i
		}
		out <- sum
	}()

	return in, out
}

func main() {
	in, out := doSomething()

	for i := 1; i <= 100; i++ {
		in <- i
	}
	close(in)

	// blocking until get the result from channel out
	res := <-out
	fmt.Println(res)
}
