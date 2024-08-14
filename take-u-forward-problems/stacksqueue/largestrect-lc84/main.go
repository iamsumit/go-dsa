package main

import "fmt"

type MyStack []int

func (m *MyStack) Push(data int) {
	*m = append(*m, data)
}

func (m *MyStack) Pop() int {
	vals := *m
	n := len(vals)
	if len(vals) == 0 {
		return -1
	}

	x := vals[n-1]
	*m = vals[0 : n-1]
	return x
}

func (m *MyStack) Top() int {
	vals := *m
	return vals[len(vals)-1]
}

func (m *MyStack) Empty() bool {
	return len(*m) == 0
}

// https://takeuforward.org/data-structure/area-of-largest-rectangle-in-histogram/
// using stack
func largestRectangleArea(heights []int) int {
	var maxArea int

	stack := MyStack{}
	for i := 0; i <= len(heights); i++ {
		// Keep looping till find the right smaller element for stack top.
		for !stack.Empty() && (i == len(heights) || heights[stack.Top()] > heights[i]) {
			height := heights[stack.Top()]
			stack.Pop()

			var width int
			if stack.Empty() {
				width = i
			} else {
				// Width will be the number of elements between left and right smaller.
				// ex: 2, 1, 5, 6, 2, 3
				// if we are at index 2 (value: 5), left smaller is at index 1 (value 1)
				// right smaller is at index 4 (value 2);
				// The gap between two is of two elements.
				// 5 is obviously the minimum element.
				// 5*2 = 10 will the area and will match this value with previously stored value.
				width = i - stack.Top() - 1
			}

			maxArea = max(maxArea, width*height)
		}

		stack.Push(i)
	}

	return maxArea
}

// Brute force
// func largestRectangleArea(heights []int) int {
// 	var maxArea int
// 	var minHeight int
// 	for i := 0; i < len(heights); i++ {
// 		maxArea = max(maxArea, heights[i])
// 		minHeight = heights[i]
// 		for j := i + 1; j < len(heights); j++ {
// 			minHeight = min(minHeight, heights[j])
// 			maxArea = max(maxArea, minHeight*(j-i+1))
// 		}
// 	}

// 	return maxArea
// }

func main() {
	fmt.Println(largestRectangleArea([]int{}), "expected: 0")
	fmt.Println(largestRectangleArea([]int{2}), "expected: 2")
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}), "expected: 10")
	fmt.Println(largestRectangleArea([]int{2, 1, 4, 6, 2, 3}), "expected: 8")
	fmt.Println(largestRectangleArea([]int{1, 2, 3, 4, 5, 6}), "expected: 12")
	fmt.Println(largestRectangleArea([]int{2, 4}), "expected: 4")
	fmt.Println(largestRectangleArea([]int{1, 1}), "expected: 2")
}
