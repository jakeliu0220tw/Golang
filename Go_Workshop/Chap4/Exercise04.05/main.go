package main

import "fmt"

func message() string {
	msg := ""
	arr := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
		msg += fmt.Sprintln(arr[i])
	}

	return msg
}

func main() {
	fmt.Println(message())
}
