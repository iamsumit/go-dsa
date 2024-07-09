package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/sudoku-solver/
// Approach 2
func solveSudoku(board [][]byte) {
	var solveFunc func() bool
	solveFunc = func() bool {
		for row := range board {
			for col := range board[row] {
				if board[row][col] == '.' {
					for val := '1'; val <= '9'; val++ {
						if isValid(board, row, col, byte(val)) {
							board[row][col] = byte(val)
							if solveFunc() {
								return true
							} else {
								board[row][col] = '.'
							}
						}
					}

					return false
				}
			}
		}

		return true
	}

	solveFunc()
}

func isValid(board [][]byte, row int, col int, val byte) bool {
	for i := 0; i < len(board); i++ {
		if (board[row][i] == val) || (board[i][col] == val) {
			return false
		}

		if board[3*(row/3)+(i/3)][3*(col/3)+i%3] == val {
			return false
		}
	}

	return true
}

// Approach 1 Simplified
// func solveSudoku(board [][]byte) {
// 	hashMap := map[int]int{}
// 	valMap := map[int]byte{}
// 	var k int
// 	for i, b := range board {
// 		for j, val := range b {
// 			if val == '.' {
// 				hashMap[k] = (i * 9) + j
// 				valMap[k] = '1'
// 				k++
// 			}
// 		}
// 	}

// 	var solveFunc func(index int)
// 	solveFunc = func(index int) {
// 		if index == len(hashMap) {
// 			return
// 		}

// 		row, col := hashMap[index]/9, hashMap[index]%9

// 		if valMap[index] > '9' {
// 			board[row][col] = '.'
// 			valMap[index] = '1'
// 			solveFunc(index - 1)
// 			return
// 		}

// 		if isValid(board, row, col, valMap[index]) {
// 			board[row][col] = valMap[index]
// 			valMap[index+1] = '1'
// 			solveFunc(index + 1)
// 			return
// 		} else {
// 			valMap[index] = valMap[index] + 1
// 			solveFunc(index)
// 			return
// 		}

// 	}

// 	solveFunc(0)
// }

// func isValid(board [][]byte, row int, col int, val byte) bool {
// 	for i := 0; i < len(board); i++ {
// 		if (board[row][i] == val) || (board[i][col] == val) {
// 			return false
// 		}
// 	}

// 	startRow, startCol := (row - (row % 3)), col-(col%3)
// 	for r := startRow; r <= startRow+2; r++ {
// 		for c := startCol; c <= startCol+2; c++ {
// 			if r == row && c == col {
// 				continue
// 			}

// 			if board[r][c] == val {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

// // Approach 1
// func solveSudoku(board [][]byte) {
// 	hashMap := map[int]int{}
// 	valMap := map[int]byte{}
// 	var k int
// 	for i, b := range board {
// 		for j, val := range b {
// 			if val == '.' {
// 				hashMap[k] = (i * 9) + j
// 				valMap[k] = '1'
// 				k++
// 			}
// 		}
// 	}

// 	var solveFunc func(cnt int, v byte)
// 	solveFunc = func(cnt int, v byte) {
// 		if cnt == len(hashMap) {
// 			return
// 		}

// 		row, col := hashMap[cnt]/9, hashMap[cnt]%9

// 		if v+1 > '9' {
// 			board[row][col] = '.'
// 			valMap[cnt] = '1'
// 			solveFunc(cnt-1, valMap[cnt-1])
// 			return
// 		}

// 		for val := v + 1; val <= '9'; val++ {
// 			if isValid(board, row, col, byte(val)) {
// 				fmt.Println(row, col)
// 				board[row][col] = byte(val)
// 				valMap[cnt] = byte(val)
// 				solveFunc(cnt+1, '0')
// 				break
// 			} else if val == '9' {
// 				board[row][col] = '.'
// 				valMap[cnt] = '1'
// 				solveFunc(cnt-1, valMap[cnt-1])
// 				break
// 			}
// 		}

// 	}

// 	solveFunc(0, '0')
// }

// func isValid(board [][]byte, row int, col int, val byte) bool {
// 	for i := 0; i < len(board); i++ {
// 		if (board[row][i] == val) || (board[i][col] == val) {
// 			return false
// 		}
// 	}

// 	startRow, startCol := (row - (row % 3)), col-(col%3)
// 	for r := startRow; r <= startRow+2; r++ {
// 		for c := startCol; c <= startCol+2; c++ {
// 			if r == row && c == col {
// 				continue
// 			}

// 			if board[r][c] == val {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

func main() {
	sudoku := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(sudoku)
	printBoard(sudoku)
}

func printBoard(sudoku [][]byte) {
	fmt.Println("  --------------------------")
	for r, s := range sudoku {
		fmt.Printf(" | ")
		for c, val := range s {
			fmt.Printf(string(val) + " ")
			if c%3 == 2 {
				fmt.Printf(" | ")
			}
		}

		fmt.Println()
		if r%3 == 2 {
			fmt.Println("  --------------------------")
		}
	}
}
