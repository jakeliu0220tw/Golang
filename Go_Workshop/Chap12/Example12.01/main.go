package main

import (
	"flag"
	"fmt"
)

func main() {
	v := flag.Int("value", -1, "a value for flag value")
	flag.Parse()
	fmt.Println(*v)
}
