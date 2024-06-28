package main

import "fmt"

// https://takeuforward.org/data-structure/count-maximum-consecutive-ones-in-the-array/
func findMaxConsecutiveOnes(nums []int) int {
	maxCount, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count = 0
		}

		maxCount = max(maxCount, count)
	}

	return maxCount
}

// func findMaxConsecutiveOnes(nums []int) int {
// 	start, end := 0, 0
// 	maxCount := 0

// 	for end < len(nums) {
// 		if nums[end] == 0 {
// 			maxCount = max(maxCount, end-start)
// 			start = end + 1
// 		}

// 		end++
// 	}

// 	return max(maxCount, end-start)
// }

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
