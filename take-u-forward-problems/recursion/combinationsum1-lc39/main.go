package main

import (
	"fmt"
	"slices"
)

// https://takeuforward.org/data-structure/combination-sum-1/
func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	subsets := [][]int{}
	sets := []int{}

	var subFunc func(index int, target int)
	subFunc = func(index, target int) {
		if index == len(candidates) {
			if target == 0 {
				subsets = append(subsets, append([]int{}, sets...))
			}

			return
		}

		if candidates[index] <= target {
			sets = append(sets, candidates[index])
			subFunc(index, target-candidates[index])
			sets = sets[:len(sets)-1]
		}

		subFunc(index+1, target)
	}

	subFunc(0, target)

	return subsets
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2}, 1))
}
