package main

import "fmt"

// https://leetcode.com/problems/next-greater-element-i/description/
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	nge := make([]int, n)
	for i := range nge {
		nge[i] = -1
	}
	st := []int{}

	for i := 2*n - 1; i >= 0; i-- {
		for len(st) > 0 && st[len(st)-1] <= nums[i%n] {
			st = st[:len(st)-1]
		}

		if i < n && len(st) > 0 {
			nge[i] = st[len(st)-1]
		}

		st = append(st, nums[i%n])
	}

	return nge
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
	fmt.Println(nextGreaterElements([]int{3, 10, 4, 2, 1, 2, 6, 1, 7, 2, 9}))
	fmt.Println(nextGreaterElements([]int{5, 7, 1, 7, 6, 0}))
}
