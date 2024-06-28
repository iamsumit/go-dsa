package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/two-sum-check-if-a-pair-with-given-sum-exists-in-array/
// Hashmap solution
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	seen[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		needle := target - nums[i]
		if val, ok := seen[needle]; ok {
			return []int{val, i}
		}

		seen[nums[i]] = i
	}

	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
