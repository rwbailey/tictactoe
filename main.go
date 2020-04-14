package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Make a board
	board := newBoard()
	// fmt.Println(board)

	// Loop forever

	// Current player = ???

	// Print current state of board
	render(board)

	// Get the move from the current board
	move := getMove()
	fmt.Printf("%T\t%v\n", move, move)

	// Make the move

	// Work out if there is a winner

	// If there is a winner, crown the champion and exit the loop

	// If there is nno winner and the board is full, declare a draw

	// Repeat until the game is over
}

type board [][]string
type move int

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
	for t == false {
		fmt.Printf("Please anter your move: ")
		m, _ := reader.ReadString('\n')
		n, t = isNumeric(strings.Trim(m, "\n"))
	}

	return move(n)
}

func isNumeric(s string) (int, bool) {
	i, err := strconv.ParseInt(s, 10, 64)
	return int(i), err == nil
}
