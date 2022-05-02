package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanUp() {
	fmt.Println("do something to cleanup ...")
}

func main() {
	// establish os channel
	sigs := make(chan os.Signal, 1)
	// register os signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer cleanUp()

	fmt.Println("program runs ...")

MyLoop:
	for {
		select {
		case <-sigs:
			fmt.Println("caught signal interrupt/terminate")
			break MyLoop
		default:
			time.Sleep(time.Microsecond * 100)
			fmt.Println("...")
		}
	}
}
