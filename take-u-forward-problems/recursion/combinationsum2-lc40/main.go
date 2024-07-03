package main

import (
	"fmt"
	"slices"
)

// https://takeuforward.org/data-structure/combination-sum-1/
func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	subsets := [][]int{}
	sets := []int{}

	var subFunc func(index int, target int)
	subFunc = func(index, target int) {
		if target == 0 {
			subsets = append(subsets, append([]int{}, sets...))
			return
		}

		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}

			if candidates[i] > target {
				break
			}

			sets = append(sets, candidates[i])
			subFunc(i+1, target-candidates[i])
			sets = sets[:len(sets)-1]
		}
	}

	subFunc(0, target)

	return subsets
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
