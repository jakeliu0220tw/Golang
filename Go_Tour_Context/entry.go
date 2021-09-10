package main

import "fmt"

func main() {
	fmt.Println("====================")
	withCancelDemo()
	fmt.Println("====================")
	withDeadlineDemo()
	fmt.Println("====================")
	withTimeoutDemo()
	fmt.Println("====================")
	withValueDemo()
}
