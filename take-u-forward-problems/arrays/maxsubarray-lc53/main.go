package main

import "fmt"

// Optimized solution
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > maxSum {
			maxSum = sum
		}

		// If sum is negative, reset it to 0
		if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}

// Brute force solution
// func maxSubArray(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	maxSum := nums[0]

// 	for i := 0; i < len(nums); i++ {
// 		sum := 0
// 		for j := i; j < len(nums); j++ {
// 			sum += nums[j]
// 			if sum > maxSum {
// 				maxSum = sum
// 			}
// 		}
// 	}

// 	return maxSum
// }

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(maxSubArray([]int{1}))                             // 1
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))                // 23
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(maxSubArray([]int{3, -2, -3, -3, 1, 3, 0}))        // 4
	fmt.Println(maxSubArray([]int{3, 1, -10, 3, 1, 3}))            // 7
	fmt.Println(maxSubArray([]int{-3, 1, -10, -3, -2, -3}))        // -1
}
