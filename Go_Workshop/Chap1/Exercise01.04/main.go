package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool
	LogLevel    = "info"
	StartupTime = time.Now()
)

func main() {
	fmt.Println("Debug =", Debug)
	fmt.Println("LogLevel =", LogLevel)
	fmt.Println("StartupTime =", StartupTime)
}
