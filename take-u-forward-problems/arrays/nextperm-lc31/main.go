package main

import (
	"fmt"
	"slices"
)

func nextPermutation(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	i := len(nums) - 2
	for i > 0 && nums[i] >= nums[i+1] {
		i--
	}

	j := len(nums) - 1
	for j > 0 && nums[j] <= nums[i] {
		j--
	}

	right := j
	left := i

	for left >= i && left < right {
		if nums[left] < nums[right] {
			nums[left], nums[right] = nums[right], nums[left]
			slices.Sort(nums[i+1:])
			return
		}

		left--
	}

	slices.Sort(nums)
}

func main() {
	// nums := []int{5, 4, 7, 5, 3, 2}
	// nums := []int{1, 5, 1}
	// nums := []int{4, 2, 0, 2, 3, 2, 0}
	// nums := []int{5, 4, 3, 2, 1}
	// nums := []int{1, 2, 3, 4, 5}
	nums := []int{1, 5}
	nextPermutation(nums)

	fmt.Println("nums", nums)
}
