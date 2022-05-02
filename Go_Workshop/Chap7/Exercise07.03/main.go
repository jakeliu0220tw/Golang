package main

import "fmt"

type record struct {
	key       string
	data      interface{}
	valueType string
}

type person struct {
	name      string
	age       int
	isMarried bool
}

type animal struct {
	name     string
	category string
}

func main() {
	m := make(map[string]interface{})

	m["person"] = person{name: "jake", age: 42, isMarried: true}
	m["animal"] = animal{name: "juice", category: "cat"}
	m["int"] = 99
	m["string"] = "test"
	m["bool"] = false

	rs := []record{}
	for k, v := range m {
		r := newRecord(k, v)
		rs = append(rs, r)
	}

	for _, v := range rs {
		fmt.Printf("key:%v\n", v.key)
		fmt.Printf("data:%v\n", v.data)
		fmt.Printf("value:%v\n", v.valueType)
	}
}

func newRecord(key string, value interface{}) record {
	r := record{}
	r.key = key
	r.data = value

	switch value.(type) {
	case int:
		r.valueType = "int"
		return r
	case string:
		r.valueType = "string"
		return r
	case bool:
		r.valueType = "bool"
		return r
	case person:
		r.valueType = "person"
		return r
	case animal:
		r.valueType = "animal"
		return r
	default:
		r.valueType = "unknown"
		return r
	}
}
