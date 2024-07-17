package main

import (
	"fmt"
	"slices"
)

// https://takeuforward.org/data-structure/aggressive-cows-detailed-solution/

// binary search
func aggressiveCows(arr []int, k int) int {
	slices.Sort(arr)
	minD, maxD := 1, (arr[len(arr)-1] - arr[0])

	for minD <= maxD {
		distance := (minD + maxD) / 2
		if countCows(arr, distance) < k {
			maxD = distance - 1
		} else {
			minD = distance + 1
		}
	}

	return maxD
}

// brute force
// func aggressiveCows(arr []int, k int) int {
// 	slices.Sort(arr)
// 	minD, maxD := math.MaxInt, (arr[len(arr)-1] - arr[0])
// 	for i := 1; i < len(arr); i++ {
// 		minD = min(minD, arr[i]-arr[i-1])
// 	}

// 	for distance := maxD; distance >= minD; distance-- {
// 		if countCows(arr, distance) >= k {
// 			return distance
// 		}
// 	}

// 	return -1
// }

func countCows(arr []int, distance int) int {
	count := 1
	sumDistance := 0
	for i := 1; i < len(arr); i++ {
		if (sumDistance + arr[i] - arr[i-1]) < distance {
			sumDistance += arr[i] - arr[i-1]
		} else {
			count++
			sumDistance = 0
		}
	}

	return count
}

func main() {
	fmt.Println(aggressiveCows([]int{0, 3, 4, 7, 10, 9}, 4))
	fmt.Println(aggressiveCows([]int{4, 2, 1, 3, 6}, 2))
}
