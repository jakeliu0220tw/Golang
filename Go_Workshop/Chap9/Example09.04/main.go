package main

import (
	"errors"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("start of app")
	if err := errors.New("aborted!"); err != nil {
		log.Println(err)
	}
	log.Println("end of app")
}
