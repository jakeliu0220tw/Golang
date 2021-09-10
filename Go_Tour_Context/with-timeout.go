package main

import (
	"context"
	"fmt"
	"time"
)

const anotherShortDuration = 100 * time.Millisecond

func withTimeoutDemo() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), anotherShortDuration)
	defer cancel()

	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println("ctx's done channel is closed")
			fmt.Println(ctx.Err()) // prints "context deadline exceeded"
			return
		}
	}
}
