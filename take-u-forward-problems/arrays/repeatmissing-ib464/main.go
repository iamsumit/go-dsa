package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/find-the-repeating-and-missing-numbers/
// with summation formula
func repeatedNumber(nums []int) []int {
	n := len(nums)

	Sn := (n * (n + 1)) / 2
	S2n := (n * (n + 1) * (2*n + 1)) / 6

	S, S2 := 0, 0
	for i := 0; i < n; i++ {
		S += nums[i]
		S2 += (nums[i] * nums[i])
	}

	// S-Sn = X-Y:
	val1 := S - Sn

	// S2-S2n = X^2-Y^2:
	val2 := S2 - S2n

	// X+Y = (X^2-Y^2)/(X-Y):
	val2 = val2 / val1

	// X = ((X+Y)+(X-Y))/2 and Y = X-(X-Y),
	// Here, X-Y = val1 and X+Y = val2...
	x := (val1 + val2) / 2
	y := x - val1

	return []int{x, y}
}

// With sorting
// func repeatedNumber(nums []int) []int {
// 	slices.Sort(nums)

// 	n := len(nums)

// 	asum := 0
// 	duplicate := 0
// 	for i := 0; i < n; i++ {
// 		asum += nums[i]

// 		if duplicate == 0 && nums[i] == nums[i+1] {
// 			duplicate = nums[i]
// 		}
// 	}

// 	esum := (n * (n + 1)) / 2

// 	return []int{
// 		duplicate,
// 		duplicate + (esum - asum),
// 	}
// }

func main() {
	fmt.Println(repeatedNumber([]int{3, 1, 2, 5, 3}))
	fmt.Println(repeatedNumber([]int{3, 1, 2, 5, 5}))
	fmt.Println(repeatedNumber([]int{3, 2, 2, 4, 5}))
	fmt.Println(repeatedNumber([]int{1, 1}))
}

// 1 + 2 + 3 + 5 + 5

// 15-14 = 3+1
// 15-16 = 5-1
