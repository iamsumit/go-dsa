package main

import "fmt"

// https://takeuforward.org/data-structure/next-greater-element-using-stack/

// Approach 2 with stack
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	set := map[int]int{}
	stack := []int{}

	for _, n := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < n {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			set[top] = n
		}
		stack = append(stack, n)
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		set[top] = -1
	}

	result := []int{}
	for _, n := range nums1 {
		result = append(result, set[n])
	}

	return result
}

// Approach 1 without stack
// func nextGreaterElement(nums1 []int, nums2 []int) []int {
// 	hm := map[int]int{}
// 	for i, val := range nums2 {
// 		hm[val] = i
// 	}

// 	result := []int{}
// 	for _, n := range nums1 {
// 		flag := false
// 		for i := hm[n] + 1; i < len(nums2); i++ {
// 			if nums2[i] > n {
// 				result = append(result, nums2[i])
// 				flag = true
// 				break
// 			}
// 		}

// 		if !flag {
// 			result = append(result, -1)
// 		}
// 	}

// 	return result
// }

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
	fmt.Println(nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}))
}
