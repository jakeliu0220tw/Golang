package main

import "fmt"

type Speaker interface {
	Speak() string
}

type person struct {
	name      string
	age       int
	isMarried bool
}

func (p person) Speak() string {
	s := "Hello, my name is "
	s += p.name
	return s
}

func (p person) String() string {
	return fmt.Sprintf("name:%v\nage:%v\nisMarried:%v", p.name, p.age, p.isMarried)
}

func main() {
	p := person{
		name:      "Jake",
		age:       42,
		isMarried: true,
	}

	fmt.Println(p.Speak())
	fmt.Println(p)
}
