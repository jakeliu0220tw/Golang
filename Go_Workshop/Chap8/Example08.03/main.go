package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	s := stringutil.Reverse("hello world")
	fmt.Println(s)
}
