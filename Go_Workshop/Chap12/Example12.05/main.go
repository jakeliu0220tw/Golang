package main

import (
	"errors"
	"fmt"
	"os"
)

func checkFile(filename string) {
	fInfo, err := os.Stat(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("file not exist: %s\n", filename)
		}
		return
	}

	fmt.Printf("name:%s\n", filename)
	fmt.Printf("size:%d\n", fInfo.Size())
	fmt.Printf("mode:%v\n", fInfo.Mode())
	fmt.Printf("modified:%v\n", fInfo.ModTime())
	fmt.Printf("isDir:%t\n", fInfo.IsDir())
}

func main() {
	checkFile("junk.txt")
	checkFile("test.txt")
}
