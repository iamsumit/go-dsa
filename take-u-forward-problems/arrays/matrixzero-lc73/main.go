package main

import "fmt"

// ## Updated better approach
func setZeroes(matrix [][]int) {
	// Used single array instead of two different arrays.
	rowColMap := [][]int{}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] != 0 {
				continue
			}

			rowColMap = append(rowColMap, []int{row, col})
		}
	}

	for _, box := range rowColMap {
		row, col := box[0], box[1]

		for i := 0; i < len(matrix); i++ {
			matrix[i][col] = 0
		}

		for j := 0; j < len(matrix[row]); j++ {
			matrix[row][j] = 0
		}
	}
}

// ## Better approach
// func setZeroes(matrix [][]int) {
// 	// Used single array instead of two different arrays.
// 	rowColMap := [][]int{}

// 	for row := 0; row < len(matrix); row++ {
// 		for col := 0; col < len(matrix[row]); col++ {
// 			if matrix[row][col] != 0 {
// 				continue
// 			}

// 			for i := 0; i < len(matrix); i++ {
// 				rowColMap = append(rowColMap, []int{i, col})
// 			}

// 			for j := 0; j < len(matrix[row]); j++ {
// 				rowColMap = append(rowColMap, []int{row, j})
// 			}
// 		}
// 	}

// 	for _, box := range rowColMap {
// 		matrix[box[0]][box[1]] = 0
// 	}
// }

// // Optimal approach
// func setZeroes(matrix [][]int) {
// 	col0 := 1

// 	for i := range matrix {
// 		for j := range matrix[i] {
// 			if matrix[i][j] == 0 {
// 				matrix[i][0] = 0

// 				if j == 0 {
// 					col0 = 0
// 				} else {
// 					matrix[0][j] = 0
// 				}
// 			}
// 		}
// 	}

// 	for i := range matrix {
// 		for j := range matrix[i] {
// 			if i == 0 || j == 0 {
// 				continue
// 			}

// 			if matrix[i][j] != 0 {
// 				if matrix[i][0] == 0 || matrix[0][j] == 0 {
// 					matrix[i][j] = 0
// 				}
// 			}
// 		}
// 	}

// 	if matrix[0][0] == 0 {
// 		for j := range matrix[0] {
// 			matrix[0][j] = 0
// 		}
// 	}

// 	if col0 == 0 {
// 		for i := range matrix {
// 			matrix[i][0] = 0
// 		}
// 	}
// }

func main() {
	fmt.Println("--------------------")
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Println("Before:")
	printMatrix(matrix)
	setZeroes(matrix)
	fmt.Println("-----")
	fmt.Println("After:")
	printMatrix(matrix)

	fmt.Println("--------------------")
	matrix = [][]int{{1, 0, 1}, {1, 1, 1}, {1, 1, 1}}
	fmt.Println("Before:")
	printMatrix(matrix)
	setZeroes(matrix)
	fmt.Println("-----")
	fmt.Println("After:")
	printMatrix(matrix)

	fmt.Println("--------------------")

	matrix = [][]int{{1, 2, 3}, {4, -1, 6}, {7, 8, 9}}
	fmt.Println("Before:")
	printMatrix(matrix)
	setZeroes(matrix)
	fmt.Println("-----")
	fmt.Println("After:")
	printMatrix(matrix)

	fmt.Println("--------------------")

	matrix = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	fmt.Println("Before:")
	printMatrix(matrix)
	setZeroes(matrix)
	fmt.Println("-----")
	fmt.Println("After:")
	printMatrix(matrix)
}

func printMatrix(matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}
