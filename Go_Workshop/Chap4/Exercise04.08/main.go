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

func getLongest(args []string) string {
	var longest string
	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}

	return longest
}

func main() {
	fmt.Println(getArgs())
	fmt.Println(getLongest(getArgs()))
}
