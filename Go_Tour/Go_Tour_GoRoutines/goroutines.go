package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// sleep 100 ms
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutinesFunc() {
	// go routine is a lightweight thread which managed by Go runtime
	// go f(x, y, z) will start a new goroutine
	go say("world")
	say("hello")
}
