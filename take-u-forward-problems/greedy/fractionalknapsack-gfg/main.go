package main

import (
	"fmt"
	"slices"
)

// https://takeuforward.org/data-structure/fractional-knapsack-problem-greedy-approach/

// accepetable by https://www.geeksforgeeks.org/problems/n-meetings-in-one-room-1587115620/1
func fractionalKnapsack(w int, arr [][]int) int {
	sum := 0

	slices.SortFunc(arr, func(a, b []int) int {
		aperc := (a[1] / a[0])
		bperc := (b[1] / b[0])
		if aperc < bperc {
			return -1
		} else if aperc > bperc {
			return 1
		}

		return 0
	})

	for _, elem := range arr {
		w = w - elem[1]
		if w >= 0 {
			sum += elem[0]
			continue
		}

		sum += (elem[0] / elem[1]) * (elem[1] + w)
		break
	}

	return sum
}

func main() {
	fmt.Println(fractionalKnapsack(50, [][]int{{100, 10}, {60, 20}, {120, 30}}))
	fmt.Println(fractionalKnapsack(50, [][]int{{100, 10}, {120, 30}, {60, 20}}))
	fmt.Println(fractionalKnapsack(21, [][]int{{8, 10}, {2, 1}, {10, 7}, {1, 7}, {9, 5}, {7, 1}, {2, 8}, {6, 6}, {4, 8}, {9, 7}}))
}
