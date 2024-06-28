package main

import "fmt"

// https://takeuforward.org/data-structure/grid-unique-paths-count-paths-from-left-top-to-the-right-bottom-of-a-matrix/
// Dynamic array
func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for idx := 0; idx < m; idx++ {
		matrix[idx] = make([]int, n)
		matrix[idx][0] = 1
	}

	for col := 0; col < n; col++ {
		matrix[0][col] = 1
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			matrix[row][col] = matrix[row-1][col] + matrix[row][col-1]
		}
	}

	return matrix[m-1][n-1]
}

// combination formula
// func uniquePaths(m int, n int) int {
// 	N := n + m - 2
// 	r := m - 1

// 	res := 1
// 	for i := 1; i <= r; i++ {
// 		res = res * (N - r + i) / i
// 	}

// 	return int(res)
// }

// // Recursive
// func uniquePaths(m int, n int) int {
// 	return countPath(0, 0, m, n)
// }

// func countPath(i, j, n, m int) int {
// 	if i == n-1 && j == m-1 {
// 		return 1
// 	} else if i >= n || j >= m {
// 		return 0
// 	} else {
// 		return countPath(i+1, j, n, m) + countPath(i, j+1, n, m)
// 	}
// }

func main() {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(23, 12))
	fmt.Println(uniquePaths(1, 1))
}
