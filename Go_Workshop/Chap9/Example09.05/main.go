package main

import (
	"errors"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Example09.05:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("start of app")
	err := errors.New("aborted!")
	if err != nil {
		logger.Fatalln("exit of app")
	}
	logger.Println("end of app")
}
