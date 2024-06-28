package main

import (
	"fmt"
	"math"
	"sort"
)

// https://takeuforward.org/data-structure/merge-overlapping-sub-intervals/
// Optimal approach
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}

		return intervals[i][1] > intervals[j][1]
	})

	index := 0

	for i := 0; i <= len(intervals)-1; i++ {
		if i < len(intervals)-1 && intervals[i][1] >= intervals[i+1][0] {
			intervals[i+1][0], intervals[i+1][1] = intervals[i][0], int(math.Max(float64(intervals[i][1]), float64(intervals[i+1][1])))
			continue
		}

		intervals[index] = intervals[i]
		index = index + 1
	}

	return intervals[:index]
}

// With append in the loop
// func merge(intervals [][]int) [][]int {
// 	overlapInterval := [][]int{}

// 	sort.Slice(intervals, func(i, j int) bool {
// 		if intervals[i][0] != intervals[j][0] {
// 			return intervals[i][0] < intervals[j][0]
// 		}

// 		return intervals[i][1] > intervals[j][1]
// 	})

// 	for i := 0; i <= len(intervals)-1; i++ {
// 		if i < len(intervals)-1 && intervals[i][1] >= intervals[i+1][0] {
// 			intervals[i+1][0], intervals[i+1][1] = intervals[i][0], int(math.Max(float64(intervals[i][1]), float64(intervals[i+1][1])))
// 			continue
// 		}

// 		overlapInterval = append(overlapInterval, []int{intervals[i][0], intervals[i][1]})
// 	}

// 	return overlapInterval
// }

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 3}, {2, 4}, {3, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 3}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 4}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 1}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 0}}))
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {1, 10}}))
	fmt.Println(merge([][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}))
}
