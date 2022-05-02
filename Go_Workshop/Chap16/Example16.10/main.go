package main

import (
	"fmt"
	"sync"
)

// this example demonstrate how encaplate worker(goroutine + channel) into struct workers
type Workers struct {
	in        chan int
	out       chan int
	workerNum int
	mtx       sync.Mutex
}

func (w *Workers) init(maxWorkers int, maxData int) {
	w.in = make(chan int, maxData)
	w.out = make(chan int, maxWorkers)
	w.mtx = sync.Mutex{}

	// init goroutine as worker
	for i := 1; i <= maxWorkers; i++ {
		// task started, increase workerNum
		w.mtx.Lock()
		w.workerNum++
		w.mtx.Unlock()

		go w.worker()
	}
}

func (w *Workers) worker() {
	sum := 0
	for i := range w.in {
		sum += i
	}
	w.out <- sum

	// task completed, reduce workerNum
	w.mtx.Lock()
	w.workerNum--
	w.mtx.Unlock()

	if w.workerNum <= 0 {
		close(w.out)
	}
}

func (w *Workers) addData(data int) {
	w.in <- data
}

func (w *Workers) getResult() int {
	close(w.in)

	res := 0
	for i := range w.out {
		res += i
	}

	return res
}

func main() {
	maxWorkers := 10
	maxData := 100
	workers := Workers{}

	workers.init(maxWorkers, maxData)
	for i := 1; i <= maxData; i++ {
		workers.addData(i)
	}

	res := workers.getResult()
	fmt.Println(res)
}
