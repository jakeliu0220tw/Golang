package main

import (
	"fmt"
	"reflect"
)

// we could use tags to set default value of struct
type Server struct {
	ServerName string `name:"WebServer"`
	ServerIP   string `ip:"127.0.0.1"`
}

func demo1() {
	s := Server{"Happy001", "192.168.0.1"}
	fmt.Println("s = ", s)
	sr := reflect.TypeOf(s)
	fmt.Println("sr =", sr)

	field1 := sr.Field(0)
	fmt.Printf("name: %v \n", field1.Tag.Get("name"))

	field2 := sr.Field(1)
	fmt.Printf("ip: %v \n", field2.Tag.Get("ip"))
}
