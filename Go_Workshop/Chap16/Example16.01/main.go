package main

import (
	"log"
	"sync"
)

func sum(begin int, end int, result *int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	for i := begin; i <= end; i++ {
		mtx.Lock()
		*result += i
		mtx.Unlock()
	}
	wg.Done()
}

func main() {
	result := 0
	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}

	wg.Add(4)
	go sum(1, 25, &result, wg, mtx)
	go sum(26, 50, &result, wg, mtx)
	go sum(51, 75, &result, wg, mtx)
	go sum(76, 100, &result, wg, mtx)
	wg.Wait()

	log.SetFlags(0)
	log.Println(result)
}
