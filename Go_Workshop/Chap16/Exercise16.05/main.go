package main

import "fmt"

func echo(ch chan string) {
	msg := <-ch
	echo := "echo >> " + msg
	ch <- echo
}

func main() {
	ch := make(chan string)
	msg := "hello"
	defer close(ch)

	go echo(ch)
	fmt.Println(msg)
	ch <- msg
	msg = <-ch
	fmt.Println(msg)
}
