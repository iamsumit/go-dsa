package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/find-the-majority-element-that-occurs-more-than-n-2-times/
// Moore's Voting Algo
func majorityElement(nums []int) int {
	var count, elem int
	for _, n := range nums {
		if count == 0 {
			count = 1
			elem = n
		} else if elem == n {
			count++
		} else {
			count--
		}
	}

	return elem
}

// With hashmap
// func majorityElement(nums []int) int {
// 	hashMap := map[int]int{}

// 	for _, num := range nums {
// 		hashMap[num]++

// 		if hashMap[num] > len(nums)/2 {
// 			return num
// 		}
// 	}

// 	return -1
// }

// With sort and counters
// func majorityElement(nums []int) int {
// 	slices.Sort(nums)
// 	var maxCount int

// 	count := 1
// 	for i, n := range nums {
// 		if i+1 < len(nums) && n == nums[i+1] {
// 			count++
// 		} else {
// 			count = 1
// 		}

// 		if count > maxCount {
// 			maxCount = count
// 		}

// 		if maxCount > len(nums)/2 {
// 			return n
// 		}
// 	}

// 	return -1
// }

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
