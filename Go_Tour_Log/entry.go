package main

// https://pkg.go.dev/log

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
}

func logDemo1() {
	log.Println("This is a demo of Println()")
}

func main() {
	fmt.Println("====================")
	logDemo1()
}
