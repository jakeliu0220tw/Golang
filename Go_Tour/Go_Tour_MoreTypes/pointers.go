package main

import "fmt"

// pointer in Go is very similar with C
func pointersFunc() {
	i := 42

	p := &i                      // pointer to i
	fmt.Printf("*p = %d \n", *p) // read i through the pointer
	*p = 21                      // set i through the pointer
	fmt.Printf("i = %d\n", i)    // see the new value of i

	j := 2701
	p = &j                      // point to j
	fmt.Printf("*p = %d\n", *p) // read j through the pointer
	*p = *p / 27                // divide j through the pointer
	fmt.Printf("j = %d\n", j)   // see the new value of j
}
