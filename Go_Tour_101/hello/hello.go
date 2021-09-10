package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of the predefined logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("logmsg: ")
	log.SetFlags(0)

	// request a greeting message.
	message, err := greetings.Hello("Jake")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	// a slice of names
	names := []string{
		"Jake", "Wells", "Edgar",
	}
	messages, err := greetings.Hellos(names)

	// if an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned map of messages to the console
	// to the console
	fmt.Println(messages)
}
