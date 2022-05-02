package main

import "fmt"

func message() string {
	arr := [...]string{"ready", "to", "go"}
	return fmt.Sprintln(arr[0], arr[1], arr[2])
}

func main() {
	fmt.Println(message())
}
