package main

import (
	"fmt"
	"runtime"
)

func switchFunc() {
	fmt.Print("Go runs on ... ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Print("MacOS. \n")
	case "linux":
		fmt.Print("Linux. \n")
	default:
		fmt.Printf("%s. \n", os)
	}
}
