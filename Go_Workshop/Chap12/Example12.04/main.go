package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("os.Create + f.Write")
	f, err := os.Create("test1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte("data from Write\n"))
	f.WriteString("data from WriteString\n")

	fmt.Println("os.WriteFile")
	err = os.WriteFile("test2.txt", []byte("data from WriteFile\n"), 0666)
	if err != nil {
		panic(err)
	}
}
