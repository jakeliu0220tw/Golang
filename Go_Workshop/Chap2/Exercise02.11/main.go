package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(8) // generate 0 to 8

		if r%3 == 0 {
			fmt.Println("skip")
			continue
		} else if r%2 == 0 {
			fmt.Println("break")
			break
		}

		fmt.Println("r =", r)
	}
}
