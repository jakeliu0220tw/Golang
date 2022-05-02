package main

import (
	"fmt"
	"time"
)

func main() {
	// Debug := false
	// LogLevel := "info"
	// StartupTime := time.Now()

	Debug, LogLevel, StartupTime := false, "info", time.Now()
	fmt.Println("Debug =", Debug)
	fmt.Println("LogLevel =", LogLevel)
	fmt.Println("StartupTime =", StartupTime)
}
