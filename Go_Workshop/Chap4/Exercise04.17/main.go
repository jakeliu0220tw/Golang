package main

import "fmt"

type user struct {
	name    string
	age     int
	balance float64
	member  bool
}

func getUsers() []user {
	var users []user

	u1 := user{
		name:    "Jake",
		age:     42,
		balance: 4000000,
		member:  true,
	}

	u2 := user{
		name:    "Eva",
		age:     49,
		balance: 4000000,
		member:  true,
	}

	u3 := user{
		name:    "Dezember",
		age:     13,
		balance: 1000,
		member:  true,
	}

	u4 := user{
		name:    "Juice",
		age:     12,
		balance: 0,
		member:  false,
	}

	users = append(users, u1, u2, u3, u4)
	return users
}

func main() {
	users := getUsers()
	for i := 0; i < len(users); i++ {
		fmt.Printf("%v: %#v\n", i, users[i])
	}
}
