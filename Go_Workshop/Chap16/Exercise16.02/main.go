package main

import (
	"fmt"
	"sync"
)

func sum(begin int, end int, sum *int, wg *sync.WaitGroup) {
	*sum = 0

	for i := begin; i <= end; i++ {
		*sum += i
	}

	if wg != nil {
		wg.Done() // report end of goroutine
	}
}

func main() {
	var s1, s2 int
	wg := &sync.WaitGroup{}

	wg.Add(2) // 2 goroutine
	go sum(1, 10, &s1, wg)
	go sum(1, 100, &s2, wg)

	wg.Wait()
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
}
