package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	res := work(4, 1, 100)
	fmt.Println("result:", res)
	time.Sleep(time.Second)
}

func work(workers int, from int, to int) int {
	out := make(chan int, workers)
	done := make(chan bool, workers)

	// create workers
	for i := 1; i <= workers; i++ {
		go worker(i, out, done)
	}

	// get random number from channel "out"
	res := 0
	for i := range out {
		if i >= 90 {
			res = i
			break
		}
	}
	defer close(done)

	return res
}

// channel "done" is responsible for "exit" signal
// worker will return random number between 1-99, and put them in channel "out"
func worker(id int, out chan int, done chan bool) {
	for {
		n := rand.Intn(100)
		select {
		case out <- n:
			fmt.Printf("id:%d => %d\n", id, n)
		case <-done:
			fmt.Printf("id:%d => close\n", id)
			return
		}
	}
}
