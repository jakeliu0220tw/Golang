package main

import "fmt"

// write a program to reverse vowels in a given input.
// The non-vowels characters should remain unchanged.
// for example:
// FACEBOOK => FOCOBAK
// APPLE => EPPLA
// vowels: A, E, I, O, U

func revertVowels(str string) string {
	if len(str) <= 1 {
		return str
	}

	uft8Str := []rune(str)
	vowels := []int{}
	// fmt.Println(str)

	// keep idx of "A,E,I,O,U"
	for i := 0; i < len(str); i++ {
		if uft8Str[i] == rune('A') || uft8Str[i] == rune('E') || uft8Str[i] == rune('I') || uft8Str[i] == rune('O') || uft8Str[i] == rune('U') {
			vowels = append(vowels, i)
		}
	}
	// fmt.Println(vowels)

	if len(vowels) == 0 {
		return str
	}

	// revert string by vowels[]
	j := len(vowels) - 1
	for i := 0; i < len(vowels); i++ {
		left := vowels[i]
		right := vowels[j]
		if i >= j {
			break
		}

		// revert
		temp := uft8Str[left]
		uft8Str[left] = uft8Str[right]
		uft8Str[right] = temp

		j--
	}
	// fmt.Println(string(uft8Str))

	return string(uft8Str)
}

func main() {
	str := "FACEBOOK"
	fmt.Println(str)
	fmt.Println(revertVowels(str))
}
