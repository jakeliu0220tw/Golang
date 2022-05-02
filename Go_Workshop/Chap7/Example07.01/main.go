package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct {
	name string
	age  int
}

func (c *cat) Speak() string {
	return "Purrr Meow"
}

func (c *cat) String() string {
	return fmt.Sprintf("(name:%s, age:%d)\n", c.name, c.age)
}

func main() {
	c := &cat{name: "juice", age: 13}
	fmt.Println(c.Speak())
	fmt.Println(c)
}
