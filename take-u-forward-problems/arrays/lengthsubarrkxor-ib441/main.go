package main

import (
	"fmt"
)

// Optimal
func solve(nums []int, b int) int {
	count := 0
	xorMap := map[int]int{}
	xorMap[0] = 1

	xorr := 0

	for i := range nums {
		xorr = xorr ^ nums[i]

		count += xorMap[xorr^b]
		xorMap[xorr]++
	}

	return count
}

// Better
// func solve(nums []int, b int) int {
// 	length := len(nums)
// 	count := 0

// 	for i := range nums {
// 		xorr := 0
// 		for j := i; j < length; j++ {
// 			xorr = xorr ^ nums[j]

// 			if xorr == b {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }

// Brute force
// func solve(nums []int, b int) int {
// 	length := len(nums)
// 	count := 0

// 	for i := range nums {
// 		for j := i; j < length; j++ {

// 			xorr := 0
// 			for k := i; k <= j; k++ {
// 				xorr = xorr ^ nums[k]
// 			}

// 			if xorr == b {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }

func main() {
	fmt.Println(solve([]int{4, 2, 2, 6, 4}, 6))
	fmt.Println(solve([]int{5, 6, 7, 8, 9}, 5))
	fmt.Println(solve([]int{4, 2, 2}, 6))
}
