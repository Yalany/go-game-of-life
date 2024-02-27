package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	gameOfLife(32, 16)
}

const (
	aliveCell = "#"
	deadCell  = "_"
)

func gameOfLife(width, height int) {
	board := make([][]bool, height)
	for i := range board {
		board[i] = make([]bool, width)
	}
	fillBoard(board)
	printBoard(board)
	for {
		time.Sleep(500 * time.Millisecond)
		board = nextState(board)
		printBoard(board)
	}
}

func fillBoard(board [][]bool) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] = rand.Int31n(2) == 0
		}
	}
}

func printBoard(board [][]bool) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] {
				fmt.Print(aliveCell)
			} else {
				fmt.Print(deadCell)
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func nextState(board [][]bool) [][]bool {
	nextState := make([][]bool, len(board))
	for i := range nextState {
		nextState[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			nextState[i][j] = getNodeNextState(board, i, j)
		}
	}
	return nextState
}

func getNodeNextState(board [][]bool, x, y int) bool {
	var neighbors int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if 0 <= x+i && x+i < len(board) && 0 <= y+j && y+j < len(board[0]) && board[x+i][y+j] {
				neighbors++
			}
		}
	}
	if board[x][y] {
		neighbors--
	}
	if !board[x][y] && neighbors == 3 {
		return true
	}
	if board[x][y] && (neighbors < 2 || neighbors > 3) {
		return false
	}
	return board[x][y]
}
