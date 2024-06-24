package main

import (
	"fmt"
	"slices"
)

// Optimal approach
func rotate(matrix [][]int) {
	for row := range matrix {
		for col := row; col < len(matrix[row]); col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}

	for row := range matrix {
		slices.Reverse(matrix[row])
	}
}

// Brute force
// func rotate(matrix [][]int) {
// 	newMatrix := make([][]int, len(matrix))
// 	for row := range matrix {
// 		newMatrix[row] = make([]int, len(matrix[0]))
// 	}

// 	for row := 0; row < len(matrix); row++ {
// 		for col := 0; col < len(matrix[row]); col++ {
// 			newMatrix[col][len(matrix)-1-row] = matrix[row][col]
// 		}
// 	}

// 	copy(matrix, newMatrix)
// }

func main() {
	fmt.Println("--------------------")
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Before:")
	printMatrix(nums)
	rotate(nums)
	fmt.Println("-----")
	fmt.Println("After:")
	printMatrix(nums)
	fmt.Println("--------------------")
}

func printMatrix(matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}
