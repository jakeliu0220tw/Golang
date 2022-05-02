package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	// Debug, LogLevel, StartupTime := getConfig()

	// var (
	// 	Debug       bool
	// 	LogLevel    string
	// 	StartupTime time.Time
	// )
	// Debug, LogLevel, StartupTime = getConfig()

	var Debug, LogLevel, StartupTime = getConfig()

	fmt.Println("Debug =", Debug)
	fmt.Println("LogLevel =", LogLevel)
	fmt.Println("StartupTime =", StartupTime)
}
