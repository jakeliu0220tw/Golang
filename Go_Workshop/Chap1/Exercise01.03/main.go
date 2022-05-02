package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool      = false
	LogLevel    string    = "info"
	StartupTime time.Time = time.Now()
)

func main() {
	fmt.Println("Debug =", Debug)
	fmt.Println("LogLevel =", LogLevel)
	fmt.Println("StartupTime =", StartupTime)
}
