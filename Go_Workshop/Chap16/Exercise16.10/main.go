package main

import (
	"errors"
	"fmt"
)

func read(in chan string, done chan bool) {
	for c := range in {
		fmt.Println(c)
	}
	done <- true
}

func main() {
	strs := []string{"a", "b", "c", "d", "e", "f", "g"}
	workers := 4
	in := make(chan string, len(strs))
	done := make(chan bool, workers)

	// initialize workers
	for i := 1; i <= workers; i++ {
		go read(in, done)
	}

	// feed data
	for _, v := range strs {
		in <- v
	}
	close(in)

	// wait until all workers completed
	for i := 1; i <= workers; i++ {
		if res := <-done; res != true {
			panic(errors.New("res should be true"))
		}
	}
	close(done)
}
