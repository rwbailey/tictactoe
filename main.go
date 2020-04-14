package main

import "fmt"

func main() {

	// Make a board
	board := newBoard()
	// fmt.Println(board)

	// Loop forever

	// Current player = ???

	// Print current state of board
	render(board)
	// Get the move from the current board

	// Make the move

	// Work out if there is a winner

	// If there is a winner, crown the champion and exit the loop

	// If there is nno winner and the board is full, declare a draw

	// Repeat until the game is over
}

type board [][]string

func newBoard() board {
	return board{
		[]string{" ", " ", " "},
		[]string{" ", " ", " "},
		[]string{" ", " ", " "}}
}

func render(b board) {

	fmt.Printf("-------\n")
	for _, row := range b {
		fmt.Printf("|")
		for _, val := range row {
			fmt.Printf("%v|", val)
		}
		fmt.Printf("\n-------\n")
	}
}
