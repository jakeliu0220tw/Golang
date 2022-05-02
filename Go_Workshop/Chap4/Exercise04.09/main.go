package main

import (
	"fmt"
	"os"
)

func getArgs() []string {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func appendArgs(args []string) []string {
	var newArgs []string
	newArgs = append(newArgs, args...)
	newArgs = append(newArgs, "i", "am", "certainly", "sure")
	return newArgs
}

func main() {
	fmt.Println(getArgs())
	fmt.Println(appendArgs(getArgs()))
}
