package main

import (
	"fmt"
	"log"
	"os"

	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

type Job struct {
	command string
	logger  *log.Logger
}

type Job2 struct {
	string
	*log.Logger
}

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(morestrings.ReverseRunes("Hello, world!"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	j := Job{"start", log.New(os.Stderr, "Job: ", log.Ldate)}
	j.logger.Println("hello job")

	j2 := Job2{"start2", log.New(os.Stderr, "Job2: ", log.Ldate)}
	j2.Println("hello job2")
}
