package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/median-of-row-wise-sorted-matrix/

// Binary search approach
func median(matrix [][]int, R, C int) int {
	low, high := matrix[0][0], matrix[0][C-1]
	for _, m := range matrix {
		low = min(low, m[0])
		high = max(high, m[C-1])
	}

	req := (R * C) / 2
	for low <= high {
		mid := (low + high) / 2
		smallEqual := countSmallEqual(matrix, C, mid)
		if smallEqual <= req {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

func countSmallEqual(matrix [][]int, C, midX int) int {
	cnt := 0
	for _, m := range matrix {
		cnt += upperBound(m, midX, C)
	}

	return cnt
}

func upperBound(row []int, midX, C int) int {
	low := 0
	high := C - 1
	ans := C

	for low <= high {
		mid := (low + high) / 2
		if row[mid] > midX {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return ans
}

// brute force approach
// func median(matrix [][]int, R, C int) int {
// 	mKey := ((R * C) / 2)

// 	if R == 1 {
// 		return matrix[0][mKey]
// 	}

// 	sortedList := []int{}
// 	for _, m := range matrix {
// 		sortedList = append(sortedList, m...)
// 	}
// 	slices.Sort(sortedList)

// 	return sortedList[mKey]
// }

func main() {
	// fmt.Println(median([][]int{{1, 4, 9}, {2, 5, 6}, {3, 8, 7}}, 3, 3))
	// fmt.Println(median([][]int{{1, 3, 8}, {2, 3, 4}, {1, 2, 5}}, 3, 3))
	// fmt.Println(median([][]int{{1, 3, 5}, {2, 6, 9}, {3, 6, 9}}, 3, 3))
	// fmt.Println(median([][]int{{1}, {2}, {3}}, 3, 1))
	// fmt.Println(median([][]int{{1, 2, 3}}, 1, 3))
	fmt.Println(median([][]int{{1, 7, 10}}, 1, 3))
}
