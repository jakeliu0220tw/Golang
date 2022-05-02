package main

import (
	"fmt"
	"sync"
)

func read(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// this loop won't stop until channel is closed
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go read(ch, wg)
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}
