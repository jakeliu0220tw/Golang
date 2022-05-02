package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	users := map[string]string{
		"101": "Jake",
		"202": "Eva",
		"303": "Dezember",
	}

	return users
}

func getUser(id string) (string, bool) {
	users := getUsers()
	name, exist := users[id]
	return name, exist
}

func main() {
	fmt.Println("Users:", getUsers())

	id := "109"
	name, exist := getUser(id)
	if !exist {
		fmt.Println("id:", id, "not exist")
		os.Exit(1)
	}
	fmt.Println("id:", id, "=> name:", name)
}
