package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(v int) {
			fmt.Println(v)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("done")
}

// output
// 3
// 5
// 9
// 7
// 8
// 4
// 1
// 0
// 2
// 6
// done
