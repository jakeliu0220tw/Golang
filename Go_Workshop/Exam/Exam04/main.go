package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case <-ch:
		fmt.Println("read")
	case ch <- 1:
		fmt.Println("write")
	}

}

// output
// fatal error: all goroutines are asleep - deadlock!
// goroutine 1 [select]:
// main.main()
//         D:/GitRepo/MyGoPractice/TheGoWorkshop/Exam/Exam04/main.go:8 +0xbc
// exit status 2

// we could not just use one channel to read & write. That makes deadlock happened!
