package main

import (
	"fmt"

	"go.uber.org/fx"
)

func demo1() {
	// create a fx application and run it.
	fmt.Println("creat a Fx application and run it.")
	fx.New().Run()
}
