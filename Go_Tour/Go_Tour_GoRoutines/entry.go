package main

import "fmt"

func main() {
	fmt.Println("====================")
	goroutinesFunc()
	fmt.Println("====================")
	channelsFunc()
	fmt.Println("====================")
	bufferedChannels()
	fmt.Println("====================")
	rangeAndClose()
	fmt.Println("====================")
	selectDemo()
	fmt.Println("====================")
	defaultSelectionDemo()
	fmt.Println("====================")
	mutexDemo()
}
