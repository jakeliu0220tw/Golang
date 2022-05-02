package main

import "fmt"

func main() {
	select {}
	fmt.Println("done")
}

// output
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [select (no cases)]:
// main.main()
//         D:/GitRepo/MyGoPractice/TheGoWorkshop/Exam/Exam03/main.go:6 +0x27
// exit status 2

// program wil wait until any case in selct{} is satisfied and executed ...
