package main

import (
	"fmt"
	"runtime"
)

func main() {
	var list []int8
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}

	// print heap memory status
	var status runtime.MemStats
	runtime.ReadMemStats(&status)
	fmt.Printf("TotalAlloc Heap = %v MB\n", status.TotalAlloc/1024/1024)
}
