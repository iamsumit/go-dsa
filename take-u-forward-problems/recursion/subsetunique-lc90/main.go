package main

import (
	"fmt"
	"slices"
)

// https://takeuforward.org/data-structure/subset-ii-print-all-the-unique-subsets/
func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)

	subsets := [][]int{}
	set := []int{}
	var subsetFunc func(index int)
	subsetFunc = func(index int) {
		subsets = append(subsets, append([]int{}, set...)) // Slices are pointers to the underline array.
		for i := index; i < len(nums); i++ {
			if i != index && nums[i] == nums[i-1] {
				continue
			}

			set = append(set, nums[i])
			subsetFunc(i + 1)
			set = set[:len(set)-1]
		}
	}

	subsetFunc(0)
	return subsets
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{0}))
	fmt.Println(subsetsWithDup([]int{1, 2, 3, 2}))
	fmt.Println(subsetsWithDup([]int{1, 2, 2, 3}))
	fmt.Println(subsetsWithDup([]int{4, 4, 4, 1, 4}))
	fmt.Println(subsetsWithDup([]int{2, 1, 2, 1, 3}))
}
