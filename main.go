package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Welcome to Go Snake!")

	// Let's start with a 2D array of the game board.
	board := make([][]rune, 10)
	for i := range board {
		board[i] = make([]rune, 10)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	// Next, initialize the snake data
	// Snake data is a list of tuples, where each tuple represents a position of the piece of the snake on the board.
	snake := make([][]int, 10*10) // Size is max size of the snake, which is the size of the board.
	snake[0] = []int{2, 3}
	snake[1] = []int{2, 2}
	snake[2] = []int{2, 1}

	// Loop until game over.
	for !IsGameOver() {
		PrintBoard(&board, &snake)
		// Move snake.
		UpdateSnake(&snake)
		// Delay.
		time.Sleep(time.Second)
	}
}

func IsGameOver() bool {
	return false
}

func PrintBoard(board *[][]rune, snake *[][]int) {
	ClearTerminal()

	for i, row := range *board {
		for j, col := range row {
			if HasSnakePiece(i, j, snake) {
				fmt.Print("*")
			} else {
				fmt.Print(string(col))
			}
		}
		fmt.Println()
	}
}

func ClearTerminal() {
	cmd := exec.Command("clear") // Linux/macOS
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func HasSnakePiece(row int, col int, snake *[][]int) bool {
	for _, piece := range *snake {
		if len(piece) > 0 && piece[0] == row && piece[1] == col {
			return true
		}
	}

	return false
}

func UpdateSnake(snake *[][]int) {
	head := (*snake)[0]
	oldRow := head[0]
	oldCol := head[1]

	head[1]++

	for _, piece := range (*snake)[1:] {
		if len(piece) > 0 {
			currRow := piece[0]
			currCol := piece[1]
			piece[0] = oldRow
			piece[1] = oldCol
			oldRow = currRow
			oldCol = currCol
		}
	}
}
