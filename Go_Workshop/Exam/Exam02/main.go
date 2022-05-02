package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	v := 0
	for i := 0; i < 10; i++ {
		v = i
		go func() {
			fmt.Println(v)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("done")
}

// we could not make sure the value of v
// output
// 9
// 3
// 9
// 9
// 9
// 9
// 9
// 9
// 9
// 9
// done
