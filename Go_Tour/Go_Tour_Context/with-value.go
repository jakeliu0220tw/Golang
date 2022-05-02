package main

import (
	"context"
	"fmt"
)

// func WithValue(parent Context, key, val interface{}) Context
// WithValue() returns a copy of parent in which the value associated with key is val
// The provided key must be comparable and should not be any other built-in type to avoid collisions between packages using context

func withValueDemo() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Printf("key: \"%v\" found value: %v\n", k, v)
			return
		}
		fmt.Println("key not found:", k)
	}

	ctx := context.WithValue(context.Background(), favContextKey("language"), "Go")

	f(ctx, favContextKey("language"))
	f(ctx, favContextKey("color"))
}
