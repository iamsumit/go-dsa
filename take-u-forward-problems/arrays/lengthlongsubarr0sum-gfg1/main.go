package main

import (
	"fmt"
	"math"
)

// https://takeuforward.org/data-structure/length-of-the-longest-subarray-with-zero-sum/
// Using hash map
func maxLenSubarray(nums []int) int {
	sumMap := map[int]int{}
	sum := 0
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == 0 {
			maxLen = i + 1
			continue
		}

		if count, ok := sumMap[sum]; ok {
			maxLen = int(math.Max(float64(maxLen), float64(i-count)))
		} else {
			sumMap[sum] = i
		}
	}

	return maxLen
}

// Using nested loop
// func maxLenSubarray(nums []int) int {
// 	maxLen := 0

// 	for i := 0; i < len(nums); i++ {
// 		sum := nums[i]
// 		count := 1

// 		right := len(nums) - 1

// 		for right >= i+1 {
// 			for j := i + 1; j <= right; j++ {
// 				sum += nums[j]
// 				count++
// 			}

// 			if sum == 0 {
// 				break
// 			} else {
// 				count = 1
// 				sum = nums[i]
// 			}

// 			right--
// 		}

// 		if sum == 0 && count > maxLen {
// 			return count
// 		}
// 	}

// 	return maxLen
// }

func main() {
	fmt.Println(maxLenSubarray([]int{9, -3, 3, -1, 6, -5}))
	// fmt.Println(maxLenSubarray([]int{9, -9, -3, 3, -1, 6, -5}))
	// fmt.Println(maxLenSubarray([]int{6, -2, 2, -8, 1, 7, 4, -10}))
	// fmt.Println(maxLenSubarray([]int{1, 0, -5}))
	// fmt.Println(maxLenSubarray([]int{1, 3, -5, 6, -2}))
}
