package main

import (
	"fmt"
	"strings"
)

func slicesOfSlice() {
	// create a tic-tac-toe board
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// the players take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	// print the board
	for i := 0; i < len(board); i++ {
		// strings.Join() will concatenates string array into one single string
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
