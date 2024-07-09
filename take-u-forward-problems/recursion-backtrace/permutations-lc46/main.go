package main

import "fmt"

// https://takeuforward.org/data-structure/print-all-permutations-of-a-string-array/
func permute(nums []int) [][]int {
	perms := [][]int{}

	var nextPermFunc func(nums []int, index int)
	nextPermFunc = func(nums []int, index int) {
		if index == len(nums) {
			perms = append(perms, append([]int{}, nums...))
			return
		}

		for i := index; i < len(nums); i++ {
			nums[i], nums[index] = nums[index], nums[i]
			nextPermFunc(nums, index+1)
			nums[index], nums[i] = nums[i], nums[index]
		}

	}

	nextPermFunc(nums, 0)

	return perms
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{1, 2, 3, 4}))
	fmt.Println(permute([]int{0, 1}))
}
