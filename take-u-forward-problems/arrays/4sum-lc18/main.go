package main

import (
	"fmt"
	"slices"
)

// https://takeuforward.org/data-structure/4-sum-find-quads-that-add-up-to-a-target-value/
func fourSum(nums []int, target int) [][]int {
	list := [][]int{}

	slices.Sort(nums)

	totalNums := len(nums)

	for i := 0; i < totalNums; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicates
		}

		for j := i + 1; j < totalNums; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue // Skip duplicates
			}

			left, right := j+1, totalNums-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					list = append(list, []int{nums[i], nums[j], nums[left], nums[right]})

					for left < right && nums[left] == nums[left+1] {
						left++
					}

					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}

	return list
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
