package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board [][]string
type move [2]int

// var theBoard board
// var currentPlayer string

func main() {

	// Make a board
	theBoard := newBoard()
	currentPlayer := "X"

	// Loop forever
	for {
		// Current player = ???

		// Print current state of board
		render(theBoard)

		// Get the move from the current board
		move := getMove()

		// Check if the move is valid for the current state of the board
		if isValidMove(theBoard, move) {
			// Make the move
			theBoard = makeMove(theBoard, move, currentPlayer)

			// Switch player
			if currentPlayer == "X" {
				currentPlayer = "O"
			} else {
				currentPlayer = "X"
			}
		} else {
			fmt.Println("That move is invalid.")
		}

		// Work out if there is a winner
		win, winner := getWinner(theBoard)

		// If there is a winner, crown the champion and exit the loop
		if win {
			render(theBoard)
			fmt.Printf("\n%v wins!\n\nWould you like to play again? (y/n) ", winner)
			if !playAgain() {
				break
			}

			// Reset the game
			theBoard = newBoard()
			currentPlayer = "X"
		} else if getDraw(theBoard) {
			render(theBoard)
			fmt.Printf("\nThe game is a draw!\n\nWould you like to play again? (y/n) ")
			if !playAgain() {
				break
			}
			// Reset the game
			theBoard = newBoard()
			currentPlayer = "X"
		}
	}

	// If there is no winner and the board is full, declare a draw

	// Repeat until the game is over
}

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

func getWinner(b board) (bool, string) {
	possiblelines := [][]string{
		// Horizontals
		[]string{b[0][0], b[0][1], b[0][2]},
		[]string{b[1][0], b[1][1], b[1][2]},
		[]string{b[2][0], b[2][1], b[2][2]},

		// Verticals
		[]string{b[0][0], b[1][0], b[2][0]},
		[]string{b[0][1], b[1][1], b[2][1]},
		[]string{b[0][2], b[1][2], b[2][2]},

		// Diagonals
		[]string{b[0][0], b[1][1], b[2][2]},
		[]string{b[2][0], b[1][1], b[0][2]},
	}

	for _, v := range possiblelines {
		if v[0] != " " && v[0] == v[1] && v[1] == v[2] {
			return true, v[0]
		}
	}
	return false, ""
}

func playAgain() bool {
	reader := bufio.NewReader(os.Stdin)
	i, _ := reader.ReadString('\n')
	pa := strings.Trim(i, "\n")

	if pa == "y" {
		return true
	}
	return false
}

func getDraw(b board) bool {
	for _, r := range b {
		for _, c := range r {
			if c == " " {
				return false
			}
		}
	}
	return true
}
