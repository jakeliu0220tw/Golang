package main

import (
	"fmt"
	"sync"
)

func worker(in chan int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for i := range in {
		sum += i
	}
	out <- sum
}

func collector(in chan int, out chan int) {
	sum := 0
	for i := range in {
		sum += i
	}
	out <- sum
}

func startTasks(workers int, from int, to int) int {
	wg := &sync.WaitGroup{}

	in := make(chan int, to-from+1)
	out := make(chan int, workers)
	res := make(chan int)

	// initialize workers
	wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go worker(in, out, wg)
	}

	// initialize collector
	go collector(out, res)

	// feed data to channel "in"
	for i := from; i <= to; i++ {
		in <- i
	}

	// close channel "in" and wait until all workers completed
	close(in)
	wg.Wait()

	// close channel "out" and wait until collector completed
	close(out)
	sum := <-res

	return sum
}

func main() {
	res := startTasks(4, 1, 100)
	fmt.Println("res =", res)
}
