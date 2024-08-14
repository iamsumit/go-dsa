package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/sliding-window-maximum/
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	result := []int{}
	deq := []int{}

	for i := 0; i < len(nums); i++ {
		// Remove indices of elements not in this window
		if len(deq) > 0 && deq[0] < i-k+1 {
			deq = deq[1:]
		}

		// Remove elements from the deque that are less than the current element
		for len(deq) > 0 && nums[deq[len(deq)-1]] < nums[i] {
			deq = deq[:len(deq)-1]
		}

		// Add current element index to the deque
		deq = append(deq, i)

		// If we have processed at least k elements, add the maximum to the result list
		if i >= k-1 {
			result = append(result, nums[deq[0]])
		}
	}

	return result
}

// type MyQueue []int

// func (m *MyQueue) Push(data int) {
// 	*m = append(*m, data)
// }

// // FIFO
// func (m *MyQueue) Pop() {
// 	vals := *m
// 	if len(vals) == 0 {
// 		return
// 	}

// 	*m = vals[1:]
// }

// func (m *MyQueue) Top() int {
// 	vals := *m
// 	return vals[0]
// }

// // with queue optimal approach 1
// func maxSlidingWindow(nums []int, k int) []int {
// 	if len(nums) <= 1 || k == 1 {
// 		return nums
// 	}

// 	myQueue := &MyQueue{}
// 	maxVal := nums[0]
// 	for i := 0; i < k; i++ {
// 		myQueue.Push(nums[i])
// 		maxVal = max(maxVal, nums[i])
// 	}

// 	result := []int{maxVal}

// 	for i := 1; i <= len(nums)-k; i++ {
// 		item := myQueue.Top()
// 		myQueue.Pop()
// 		myQueue.Push(nums[i+k-1])

// 		if item != maxVal {
// 			maxVal = max(maxVal, nums[i+k-1])
// 		} else if item <= nums[i+k-1] {
// 			maxVal = nums[i+k-1]
// 		} else {
// 			// Override
// 			maxVal = nums[i]
// 			for j := i + 1; j < i+k; j++ {
// 				maxVal = max(maxVal, nums[j])
// 			}
// 		}

// 		result = append(result, maxVal)
// 	}

// 	return result
// }

// // Brute force
// func maxSlidingWindow(nums []int, k int) []int {
// 	if len(nums) <= 1 || k == 1 {
// 		return nums
// 	}

// 	result := []int{}
// 	for i := 0; i <= len(nums)-k; i++ {
// 		maxVal := nums[i]
// 		for j := i; j < i+k; j++ {
// 			maxVal = max(maxVal, nums[j])
// 		}

// 		result = append(result, maxVal)
// 	}

// 	return result
// }

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1}, 1))
	fmt.Println(maxSlidingWindow([]int{4, 0, -1, 3, 5, 3, 6, 8}, 3))
	fmt.Println(maxSlidingWindow([]int{20, 25}, 2))
	fmt.Println(maxSlidingWindow([]int{9, 8, 9, 8}, 3))
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2}, 2))
}
