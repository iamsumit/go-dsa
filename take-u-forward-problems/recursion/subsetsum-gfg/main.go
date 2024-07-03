package main

import (
	"fmt"
	"slices"
)

// https://takeuforward.org/data-structure/subset-sum-sum-of-all-subsets/

// Apprach 2
func subsetSums(nums []int) []int {
	sums := []int{}
	var recurSum func(int, int)
	recurSum = func(index, sum int) {
		if index == len(nums) {
			sums = append(sums, sum)
			return
		}

		recurSum(index+1, sum+nums[index])
		recurSum(index+1, sum)
	}

	recurSum(0, 0)
	slices.Sort(sums)
	return sums
}

// Approach 1
// func subsetSums(nums []int) []int {
// 	return recurSum(nums, []int{0})
// }

// func recurSum(nums []int, sums []int) []int {
// 	if len(nums) == 0 {
// 		slices.Sort(sums)
// 		return sums
// 	}

// 	for _, s1 := range sums[1:] {
// 		sums = append(sums, s1+nums[0])
// 	}

// 	sums = append(sums, nums[0])

// 	return recurSum(nums[1:], sums)
// }

func main() {
	fmt.Println(subsetSums([]int{5, 2, 1}))
	fmt.Println(subsetSums([]int{3, 1, 2}))
}
