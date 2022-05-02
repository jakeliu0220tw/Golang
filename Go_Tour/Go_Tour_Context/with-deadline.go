package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 100 * time.Millisecond

func withDeadlineDemo() {
	d := time.Now().Add(shortDuration)
	// func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
	// The returned context's Done channel is closed when the deadline expires
	// Or when the returned cancel function is called.
	// Or when the parent context's Done channel is closed
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println("ctx's done channel is closed")
			fmt.Println(ctx.Err())
			return
		}
	}

}
