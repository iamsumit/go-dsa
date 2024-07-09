package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/n-queen-problem-return-all-distinct-solutions-to-the-n-queens-puzzle/

// approach 3
func solveNQueens(n int) [][]string {
	result := [][]string{}

	board := make([]string, n)
	for i := range board {
		for j := 0; j < n; j++ {
			board[i] += "."
		}
	}

	leftRow := make([]int, n)
	upperDiagonal := make([]int, (2*n)-1)
	lowerDiagonal := make([]int, (2*n)-1)

	var solveFunc func(col, n int)
	solveFunc = func(col, n int) {
		if col == n {
			result = append(result, append([]string{}, board...))
			return
		}

		for row := 0; row < n; row++ {
			if leftRow[row] == 0 && lowerDiagonal[row+col] == 0 && upperDiagonal[n-1+col-row] == 0 {
				board[row] = board[row][:col] + "Q" + board[row][col+1:]
				leftRow[row] = 1
				lowerDiagonal[row+col] = 1
				upperDiagonal[n-1+col-row] = 1
				solveFunc(col+1, n)
				board[row] = board[row][:col] + "." + board[row][col+1:]
				leftRow[row] = 0
				lowerDiagonal[row+col] = 0
				upperDiagonal[n-1+col-row] = 0
			}
		}
	}

	solveFunc(0, n)

	return result
}

// appoach 2
// func solveNQueens(n int) [][]string {
// 	board := make([]string, n)
// 	for i := range board {
// 		for j := 0; j < n; j++ {
// 			board[i] += "."
// 		}
// 	}

// 	result := [][]string{}
// 	var solveFunc func(col, n int)
// 	solveFunc = func(col, n int) {
// 		if col == n {
// 			result = append(result, append([]string{}, board...))
// 			return
// 		}

// 		for row := 0; row < n; row++ {
// 			if isSafe(row, col, board, n) {
// 				board[row] = board[row][:col] + "Q" + board[row][col+1:]
// 				solveFunc(col+1, n)
// 				board[row] = board[row][:col] + "." + board[row][col+1:]
// 			}
// 		}
// 	}

// 	solveFunc(0, n)
// 	return result
// }

// func isSafe(row int, col int, board []string, n int) bool {
// 	duprow, dupcol := row, col
// 	for row >= 0 && col >= 0 {
// 		if board[row][col] == 'Q' {
// 			return false
// 		}

// 		row--
// 		col--
// 	}

// 	col, row = dupcol, duprow
// 	for col >= 0 {
// 		if board[row][col] == 'Q' {
// 			return false
// 		}

// 		col--
// 	}

// 	col, row = dupcol, duprow
// 	for row < n && col >= 0 {
// 		if board[row][col] == 'Q' {
// 			return false
// 		}

// 		row++
// 		col--
// 	}

// 	return true
// }

// Approach one
// func solveNQueens(n int) [][]string {
// 	if n == 1 {
// 		return [][]string{{"Q"}}
// 	}

// 	queens := make([][]string, n)
// 	for i := range queens {
// 		queens[i] = make([]string, n)
// 	}

// 	count := 0
// 	hashMap := map[int]int{}

// 	result := [][]string{}

// 	var setQueens func(row, col int)
// 	setQueens = func(row, col int) {
// 		if count == n {
// 			arr := []string{}
// 			for i := range queens {
// 				arr = append(arr, strings.Join(queens[i], ""))
// 			}

// 			result = append(result, append([]string{}, arr...))
// 		}

// 		if row == n {
// 			count = row - 1
// 			setQueens(row-1, hashMap[row-1]+1)
// 			return
// 		}

// 		if row == 0 && hashMap[0] == n-1 {
// 			return
// 		}

// 		for c := 0; c < n; c++ {
// 			if c == col {
// 				continue
// 			}

// 			queens[row][c] = "."
// 		}

// 		if row > 0 && col >= n {
// 			setQueens(row-1, hashMap[row-1]+1)
// 			return
// 		}

// 		canFit := true

// 		r, c := row, col
// 		for r >= 0 {
// 			if queens[r][c] == "Q" {
// 				canFit = false
// 				break
// 			}

// 			r--
// 		}

// 		r, c = row, col
// 		for r >= 0 && c >= 0 {
// 			if queens[r][c] == "Q" {
// 				canFit = false
// 				break
// 			}

// 			r--
// 			c--
// 		}

// 		r, c = row, col
// 		for r >= 0 && c < n {
// 			if queens[r][c] == "Q" {
// 				canFit = false
// 				break
// 			}

// 			r--
// 			c++
// 		}

// 		if !canFit {
// 			setQueens(row, col+1)
// 		} else {
// 			hashMap[row] = col
// 			queens[row][col] = "Q"
// 			count = row + 1
// 			setQueens(row+1, 0)
// 		}
// 	}

// 	setQueens(0, 0)

// 	return result
// }

func main() {
	// fmt.Println(solveNQueens(1))
	// fmt.Println(solveNQueens(2))
	// fmt.Println(solveNQueens(3))
	fmt.Println(solveNQueens(4))
	// fmt.Println(solveNQueens(5))
	// fmt.Println(solveNQueens(6))
	// fmt.Println(solveNQueens(7))
	// fmt.Println(solveNQueens(8))
	// fmt.Println(solveNQueens(9))
}
