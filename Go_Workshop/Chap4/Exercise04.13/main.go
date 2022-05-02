package main

import "fmt"

func getUsers() map[string]string {
	users := map[string]string{
		"101": "Eva",
		"202": "Jake",
		"303": "Dezember",
	}
	users["404"] = "MilkTea"

	return users
}

func main() {
	users := getUsers()
	fmt.Println("Users: ", users)

	for k, v := range users {
		fmt.Println("key:", k, "value:", v)
	}
}
