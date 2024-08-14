package main

import (
	"fmt"
	"math"
)

// https://www.naukri.com/code360/problems/max-of-min_982935

// Approach optimal
func maxMinWindow(nums []int) []int {
	n := len(nums)

	fmt.Println("nums", nums)

	s := []int{}
	left := make([]int, n+1)
	right := make([]int, n+1)

	for i := range left {
		left[i] = -1
	}

	for i := range right {
		right[i] = n
	}

	for i := 0; i < n; i++ {
		for len(s) > 0 && nums[s[len(s)-1]] >= nums[i] {
			s = s[:len(s)-1]
		}

		if len(s) != 0 {
			left[i] = s[len(s)-1]
		}

		s = append(s, i)
	}

	s = []int{}

	for i := n - 1; i >= 0; i-- {
		for len(s) > 0 && nums[s[len(s)-1]] >= nums[i] {
			s = s[:len(s)-1]
		}

		if len(s) != 0 {
			right[i] = s[len(s)-1]
		}

		s = append(s, i)
	}

	ans := make([]int, n+1)
	for i := range right {
		ans[i] = math.MinInt
	}

	for i := 0; i < n; i++ {
		len := right[i] - left[i] - 1
		ans[len] = max(ans[len], nums[i])
	}

	for i := n - 1; i >= 0; i-- {
		ans[i] = max(ans[i], ans[i+1])
	}

	return ans[1:]
}

// func maxMinWindow(nums []int) []int {
// 	n := len(nums)

// 	result := []int{}
// 	for k := 1; k <= n; k++ {
// 		result = append(result, maxOfMin(nums, k))
// 	}

// 	return result
// }

// func maxOfMin(nums []int, k int) int {
// 	maximum := math.MinInt

// 	deque := []int{}
// 	for i := 0; i < len(nums); i++ {
// 		if len(deque) > 0 && deque[0] < i-k+1 {
// 			deque = deque[:len(deque)-1]
// 		}

// 		for len(deque) > 0 && nums[deque[len(deque)-1]] > nums[i] {
// 			deque = deque[:len(deque)-1]
// 		}

// 		deque = append(deque, i)

// 		if i >= k-1 {
// 			maximum = max(maximum, nums[deque[0]])
// 		}
// 	}

// 	return maximum
// }

func main() {
	// fmt.Println(maxMinWindow([]int{3, 3, 4, 2, 4}))
	// fmt.Println(maxMinWindow([]int{45, -2, 42, 5, -11}))
	// fmt.Println(maxMinWindow([]int{4, 3, 2, 1}))
	fmt.Println(maxMinWindow([]int{10, 20, 30, 50, 10, 70, 30}))
}
