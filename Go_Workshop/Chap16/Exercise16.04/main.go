package main

import "fmt"

func greet(ch chan string) {
	ch <- "Hello Buddy!"
}

func main() {
	ch := make(chan string)
	defer close(ch)

	go greet(ch)
	fmt.Println(<-ch)
}
