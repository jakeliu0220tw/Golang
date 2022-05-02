package main

import (
	"fmt"
	"unicode"
)

func passwordChecker(pwd string) bool {
	pwdR := []rune(pwd) // transfer pwd into rune type

	if len(pwdR) < 8 {
		return false
	}

	hasLower := false
	hasUpper := false
	hasNumber := false
	hasSymbol := false

	for _, val := range pwdR {
		if unicode.IsUpper(val) {
			hasUpper = true
		}
		if unicode.IsLower(val) {
			hasLower = true
		}
		if unicode.IsNumber(val) {
			hasNumber = true
		}
		if unicode.IsPunct(val) {
			hasSymbol = true
		}
		if unicode.IsSymbol(val) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	fmt.Println("pwd = ''")
	if passwordChecker("") {
		fmt.Println("well format")
	} else {
		fmt.Println("invalid format")
	}

	fmt.Println("pwd = This!ISA")
	if passwordChecker("This!ISA") {
		fmt.Println("well format")
	} else {
		fmt.Println("invalid format")
	}

	fmt.Println("pwd = M@cd0nald")
	if passwordChecker("M@cd0nald") {
		fmt.Println("well format")
	} else {
		fmt.Println("invalid format")
	}
}
