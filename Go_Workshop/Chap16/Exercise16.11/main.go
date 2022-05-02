package main

import (
	"context"
	"fmt"
	"time"
)

// countNumber is some kind of producer
func countNumber(ctx context.Context, out chan int) {
	sum := 0
	for {
		select {
		case <-ctx.Done():
			out <- sum
			return
		default:
			sum += 1
			time.Sleep(time.Nanosecond)
		}
	}
}

func main() {
	out := make(chan int)

	// create a dummy context
	c := context.TODO()
	ctx, cancel := context.WithCancel(c)

	go countNumber(ctx, out)

	time.Sleep(time.Second)
	cancel()

	res := <-out
	fmt.Println(res)
}
