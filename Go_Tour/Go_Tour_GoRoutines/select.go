package main

import "fmt"

func fibonacci2(c, q chan int) {
	fmt.Println("start fibonacci2()")
	x, y := 0, 1

	// a select blocks until one of its cases can run ...
	for {
		select {
		case c <- x:
			x, y = y, (x + y)
			// fmt.Printf("x = %v, y = %v\n", x, y)

		// triggered if q is not empty ...
		case <-q:
			fmt.Println("quit")
			return
		}
	}
}

func selectDemo() {
	// create channels ... c & q
	c := make(chan int)
	q := make(chan int)

	// create a goroutine and run it
	go func() {
		// print channel c 10 times
		// channel c must have value stored in, otherwise Println() will block ...
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		q <- 0
	}()
	fibonacci2(c, q)
}
