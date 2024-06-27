package main

import (
	"fmt"
	"slices"
)

// two pointer approach with sorting
func threeSum(nums []int) [][]int {
	results := [][]int{}
	slices.Sort(nums)

	length := len(nums)
	for i := 0; i < length-2; i++ {
		left := i + 1

		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		right := length - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				results = append(results, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return results
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-1, 0, 1, 0}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}
