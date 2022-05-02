package main

import (
	"bytes"
	"log"
	"testing"
)

func Test_Main(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var s bytes.Buffer // create an io buffer
		log.SetOutput(&s)  // redirect output to this io buffer
		log.SetFlags(0)
		main()

		if s.String() != "5050\n" {
			t.Fail()
		}
	}
}
