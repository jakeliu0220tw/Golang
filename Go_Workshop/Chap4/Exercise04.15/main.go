package main

import (
	"fmt"
	"os"
)

var users = map[string]string{
	"101": "Jake",
	"202": "Eva",
	"303": "Dezember",
}

func getUsers() map[string]string {
	return users
}

func delUser(userId string) {
	delete(getUsers(), userId)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please key in the user to be deleted.")
		os.Exit(1)
	}

	userId := os.Args[1]
	fmt.Println(getUsers())
	delUser(userId)
	fmt.Println(getUsers())
}
