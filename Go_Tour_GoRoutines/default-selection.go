package main

import (
	"fmt"
	"time"
)

func defaultSelectionDemo() {
	// time.Tick() will send the current time to a ticking channel periodically ...
	chTick := time.Tick(100 * time.Millisecond)
	// time.After() will send the current time on the return channel if elapse the duration ...
	chBoom := time.After(500 * time.Millisecond)

	// "select case" is designed for dealing with channel values
	// "select case" only follow channel value, otherwise Error will rise
	for {
		select {
		case <-chTick:
			fmt.Println("tick.")
		case <-chBoom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
