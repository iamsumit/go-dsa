package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/search-in-a-sorted-2d-matrix/
// Optimal approach - binary search
func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])

	low := 0
	high := n*m - 1

	count := 1
	for low <= high {
		count++
		mid := (low + high) / 2
		row := mid / m
		col := mid % m

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

// func searchMatrix(matrix [][]int, target int) bool {
// 	row := -1

// 	left, right := 0, len(matrix)-1

// 	for left <= right {
// 		mid := left + (right-left)/2
// 		if matrix[mid][0] <= target && target <= matrix[mid][len(matrix[mid])-1] {
// 			row = mid
// 			break
// 		} else if matrix[mid][0] < target {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}

// 	if row == -1 {
// 		return false
// 	}

// 	return ColSearch(matrix[row], target)
// }

// func ColSearch(nums []int, target int) bool {
// 	left, right := 0, len(nums)-1

// 	for left <= right {
// 		mid := left + (right-left)/2

// 		if nums[mid] == target {
// 			return true
// 		} else if nums[mid] < target {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}

// 	return false
// }

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 16))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 26))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 34))
}
