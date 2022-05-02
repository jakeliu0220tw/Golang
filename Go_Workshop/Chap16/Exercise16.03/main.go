package main

import (
	"log"
	"sync"
	"sync/atomic"
)

// atomic operation (only for integer type)
func sum(begin int, end int, wg *sync.WaitGroup, res *int32) {
	for i := begin; i <= end; i++ {
		atomic.AddInt32(res, int32(i))
	}

	if wg != nil {
		wg.Done()
	}
}

func main() {
	var res int32
	wg := &sync.WaitGroup{}

	wg.Add(4)
	go sum(1, 25, wg, &res)
	go sum(26, 50, wg, &res)
	go sum(51, 75, wg, &res)
	go sum(76, 100, wg, &res)
	wg.Wait()

	log.SetFlags(0)
	log.Println(res)
}
