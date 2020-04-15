package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var theBoard board
var currentPlayer string

func main() {

	// Make a board
	theBoard = newBoard()
	currentPlayer = "X"

	// Loop forever
	for {
		// Current player = ???

		// Print current state of board
		render(theBoard)

		// Get the move from the current board
		move := getMove()

		if isValidMove(theBoard, move) {
			// Make the move
			theBoard = makeMove(theBoard, move, currentPlayer)

			if currentPlayer == "X" {
				currentPlayer = "O"
			} else {
				currentPlayer = "X"
			}
		} else {
			fmt.Println("That move is invalid.")
		}
	}

	// Work out if there is a winner

	// If there is a winner, crown the champion and exit the loop

	// If there is nno winner and the board is full, declare a draw

	// Repeat until the game is over
}

type board [][]string
type move [2]int

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

func getMove() move {
	reader := bufio.NewReader(os.Stdin)
	t := false
	var n int
	var m move

	for t == false {
		fmt.Printf("Please anter your move: ")
		i, _ := reader.ReadString('\n')
		n, t = isNumeric(strings.Trim(i, "\n"))
		if n > 9 || n < 1 {
			t = false
		}
	}

	switch {
	case n == 1:
		m = [2]int{0, 0}
	case n == 2:
		m = [2]int{0, 1}
	case n == 3:
		m = [2]int{0, 2}
	case n == 4:
		m = [2]int{1, 0}
	case n == 5:
		m = [2]int{1, 1}
	case n == 6:
		m = [2]int{1, 2}
	case n == 7:
		m = [2]int{2, 0}
	case n == 8:
		m = [2]int{2, 1}
	case n == 9:
		m = [2]int{2, 2}
	}

	return m
}

func isNumeric(s string) (int, bool) {
	i, err := strconv.ParseInt(s, 10, 64)
	return int(i), err == nil
}

func isValidMove(b board, m move) bool {
	if b[m[0]][m[1]] != " " {
		return false
	}
	return true
}

func makeMove(b board, m move, p string) board {
	n := make(board, len(b))
	copy(n, b)
	n[m[0]][m[1]] = p
	return n
}

// func getWinner() (bool, string) {

// }
