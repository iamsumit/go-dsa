package main

import "fmt"

// https://takeuforward.org/data-structure/program-to-generate-pascals-triangle/
// Solution 3 -- from Take U Forward
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		val := 1
		valRow := make([]int, i)
		valRow[0] = val

		for j := 1; j < i; j++ {
			val = (val * (i - j)) / j
			valRow[j] = val
		}

		res[i-1] = valRow
	}

	return res
}

// Solution 2 - Optimized
// func generate(numRows int) [][]int {
// 	res := make([][]int, numRows)
// 	for i := 0; i < numRows; i++ {
// 		res[i] = make([]int, i+1)
// 		res[i][0], res[i][i] = 1, 1

// 		for j := 1; j <= i-1; j++ {
// 			res[i][j] = res[i-1][j-1] + res[i-1][j]
// 		}
// 	}

// 	return res
// }

// Solution 1
// func generate(numRows int) [][]int {
// 	var res [][]int
// 	for i := 0; i < numRows; i++ {
// 		var row []int
// 		for j := 0; j <= i; j++ {
// 			if j == 0 || j == i {
// 				row = append(row, 1)
// 			} else {
// 				row = append(row, res[i-1][j-1]+res[i-1][j])
// 			}
// 		}

// 		res = append(res, row)
// 	}

// 	return res
// }

func main() {
	fmt.Println(generate(1))
	fmt.Println(generate(2))
	fmt.Println(generate(3))
	fmt.Println(generate(5))
	fmt.Println(generate(8))
}
