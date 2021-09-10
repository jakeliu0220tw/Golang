package main

// https://pkg.go.dev/gopkg.in/yaml.v2

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// Marshal serializes the value, from obj to []bytes ... encodes obj to byte array
func marshalDemo() {
	type T struct {
		F int `yaml:"a,omitempty"`
		B int
	}

	out, err := yaml.Marshal(&T{B: 2}) // Returns "b: 2\n"
	if err != nil {
		panic("oops")
	}
	fmt.Println(string(out))

	out2, err2 := yaml.Marshal(&T{F: 1}) // Returns "a: 1\nb: 0\n"
	if err2 != nil {
		panic("oops2")
	}
	fmt.Println(string(out2))

	out3, err3 := yaml.Marshal(&T{F: 777, B: 999})
	if err3 != nil {
		panic("oops3")
	}
	fmt.Println(string(out3))
}

// Unmarshal decodes []byte and assign decoded value to obj ... decodes byte array to obj
func unmarshalDemo() {
	type T struct {
		F int `yaml:"a,omitempty"`
		B int
	}

	var t T
	yaml.Unmarshal([]byte("a: 1\nb: 2"), &t)
	fmt.Println("t =", t)

	var tt T
	yaml.Unmarshal([]byte("a: 777\nb: 999"), &tt)
	fmt.Println("tt =", tt)
}

func main() {
	fmt.Println("====================")
	marshalDemo()
	fmt.Println("====================")
	unmarshalDemo()
}
