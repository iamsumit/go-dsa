package main

import (
	"fmt"
	"slices"
)

// With sorting
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slices.Sort(nums)

	maxLen, count := 1, 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] == 1 {
			count++
		} else if nums[i+1] != nums[i] {
			count = 1
		}

		if maxLen < count {
			maxLen = count
		}
	}

	return maxLen
}

// Using hashmap
// func longestConsecutive(nums []int) int {
// 	hashMap := map[int]bool{}
// 	for i := 0; i < len(nums); i++ {
// 		hashMap[nums[i]] = true
// 	}

// 	maxLen, count := 0, 0
// 	for _, n := range nums {
// 		_, ok := hashMap[n-1]
// 		if !ok {
// 			count = 1
// 			for {
// 				ok := hashMap[n+1]
// 				if !ok {
// 					break
// 				}

// 				count++
// 				n = n + 1
// 			}
// 		}

// 		if maxLen < count {
// 			maxLen = count
// 		}
// 	}

// 	return maxLen
// }

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{1, 2, 0, 1}))
}

// 100 - 100
// 4 - 4
// 200 - 200
// 1 - 2
// 3 - 4
// 2 - 3

// 3 = 4
// 7 = 8
// 2 = 3
// 5 = 6
// 8 = 8
// 4 = 5
// 6 = 7
// 0 = 1
// 1 = 2
