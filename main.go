package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:] // Grab arguments

	// Create slice for board
	var board [9][9]int

	// Solusion board
	var solutionBoard [9][9]int
	if len(args) == 1 {
		input := args[0]

		// Remove the extra double quotes from both ends
		args = strings.Split(input, " ")

	}
	// Populate boards with inputs
	for i := 0; i < len(args); i++ {
		rowInputCount := 0
		for j := 0; j < len(args[i]); j++ {
			if args[i][j] == '.' {
				board[i][j] = 0
			} else if args[i][j] >= '0' && '9' >= args[i][j] {
				board[i][j] = int(args[i][j] - '0')
			} else {
				fmt.Println("Error") // If inputs has unexpected character return "Error".
				return
			}
			rowInputCount++
		}

		if rowInputCount != 9 {
			fmt.Println("Error")
			return
		}
	}

	// Check if given input are valid
	if !firstCheck(board) {
		fmt.Println("Error")
		return
	}

	// If there is no enough arguments return "Error"
	if len(args) != 9 {
		fmt.Println("Error")
		return
	}
	solCount := 0 // Counter for multiple solusion puzzles.
	solCount = solving(board, &solutionBoard, 0)
	if solCount == 1 {
		printBoard(solutionBoard)
	} else {
		fmt.Println("Error") // If puzzle has 0 or more than 1 solusion return "Error".
	}
}

/*
	This function checks the initial board and verifies that each row, column, and 3x3 square does not have any repeated numbers.

It returns false if any number appears more than once in a row, column, or 3x3 square.
*/
func firstCheck(board [9][9]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != 0 {
				num := board[row][col]
				// Row and Column check.
				encounterInRowOrCol := 0
				for i := 0; i < 9; i++ {
					if board[row][i] == num || board[i][col] == num {
						if encounterInRowOrCol > 1 {
							return false
						}
						encounterInRowOrCol++
					}
				}
				// 3x3 square check
				rowOffset := row / 3 * 3
				colOffset := col / 3 * 3
				encounterInSquare := 0
				for i := rowOffset; i < 3; i++ {
					for j := colOffset; j < 3; j++ {
						if board[i][j] == num {
							if encounterInSquare > 1 {
								fmt.Println(num)
								return false
							}
							encounterInSquare++
						}
					}
				}
			}
		}
	}
	return true
}

// Test all possible solution remember first solution found
func solving(board [9][9]int, solutionBoard *[9][9]int, solCount int) int {
	row, col := findNextBlank(board) // find next blank spot

	// if no blank space found, treat it as solve
	if row == -1 && col == -1 {
		// first solution found
		if solCount == 0 {
			*solutionBoard = board // store solution
		}
		solCount++      // increase solution count
		return solCount // return new solution count
	}

	// If a blank spot is found, try placing numbers 1 to 9
	for num := 1; num <= 9 && solCount <= 1; num++ {
		if placeable(board, row, col, num) { 				   // check if placeable
			board[row][col] = num                              // Place the number on the board
			solCount = solving(board, solutionBoard, solCount) // Try solving the next blank space
			board[row][col] = 0                                // Unassign the number (backtrack)
		}
	}
	return solCount // Return the current solution count
}

// check if current number is placeable in current square
func placeable(board [9][9]int, row, col, num int) bool {
	// Row and column validation
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false // return false if number found in row or coumn
		}
	}
	// 3x3 square validation
	startRow := row / 3 * 3
	startCol := col / 3 * 3

	for row := startRow; row < startRow+3; row++ {
		for col := startCol; col < startCol+3; col++ {
			if board[row][col] == num {
				return false // return false if number foundn in 3x3 box
			}
		}
	}
	return true // return true if placeable
}

// find next blank space in current board
func findNextBlank(board [9][9]int) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j // return blabk space coordinates
			}
		}
	}

	// return -1, -1 if no blank space found. Treat as solved
	return -1, -1
}

// just print board
func printBoard(board [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(board[i][j])
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
