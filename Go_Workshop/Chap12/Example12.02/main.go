package main

import (
	"flag"
	"fmt"
	"os"
)

// we use flag package to parse passing parameters
func main() {
	n := flag.String("name", "", "your name")
	i := flag.Int("age", 0, "your age")
	b := flag.Bool("isMarried", false, "isMarried")
	flag.Parse()

	// check parameters
	if *n == "" {
		fmt.Println("name is necessary")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("name:", *n)
	fmt.Println("age:", *i)
	fmt.Println("isMarried:", *b)
}
