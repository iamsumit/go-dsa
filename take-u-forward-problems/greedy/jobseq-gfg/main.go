package main

import (
	"fmt"
	"slices"
)

// https://takeuforward.org/data-structure/job-sequencing-problem/

func JobScheduling(arr [][]int) []int {
	slices.SortFunc(arr, func(a, b []int) int {
		return b[2] - a[2]
	})

	maxJobs := arr[0][2]
	for _, job := range arr {
		maxJobs = max(maxJobs, job[1])
	}

	countJobs := 0
	profit := 0

	slot := make([]bool, maxJobs)

	for _, job := range arr {
		for j := (job[1]); j > 0; j-- {
			if !slot[j-1] {
				slot[j-1] = true
				countJobs += 1
				profit += job[2]
				break
			}
		}
	}

	return []int{countJobs, profit}
}

func main() {
	fmt.Println(JobScheduling([][]int{{1, 4, 20}, {2, 1, 10}, {3, 1, 40}, {4, 1, 30}}))
	fmt.Println(JobScheduling([][]int{{1, 2, 100}, {2, 1, 19}, {3, 2, 27}, {4, 1, 25}, {5, 1, 15}}))
	fmt.Println(JobScheduling([][]int{{1, 3, 288}, {2, 2, 435}, {3, 10, 401}, {4, 16, 368}, {5, 16, 248}, {6, 1, 361}, {7, 11, 108}, {8, 10, 167}, {9, 5, 251}, {10, 3, 170}, {11, 14, 156}, {12, 6, 184}, {13, 4, 370}, {14, 5, 424}, {15, 8, 397}, {16, 5, 375}, {17, 5, 218}}))
}
