package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/trapping-rainwater/
// optimal approach 2
func trap(height []int) int {
	left, right := 0, len(height)-1
	count := 0
	maxLeft, maxRight := height[left], height[right]
	for left < right {
		if maxLeft < maxRight {
			left++
			maxLeft = max(maxLeft, height[left])
			count += maxLeft - height[left]
		} else {
			right--
			maxRight = max(maxRight, height[right])
			count += maxRight - height[right]
		}
	}

	return count
}

// optimal approach 1
// func trap(height []int) int {
// 	left, right := 0, len(height)-1
// 	count := 0
// 	maxLeft, maxRight := 0, 0
// 	for left <= right {
// 		if height[left] <= height[right] {
// 			if height[left] >= maxLeft {
// 				maxLeft = height[left]
// 			} else {
// 				count += maxLeft - height[left]
// 			}

// 			left++
// 		} else {
// 			if height[right] >= maxRight {
// 				maxRight = height[right]
// 			} else {
// 				count += maxRight - height[right]
// 			}

// 			right--
// 		}
// 	}

// 	return count
// }

// with two pointer approach with storing last maxt height and skip
// better than brute force
// func trap(height []int) int {
// 	left, right := 0, len(height)-1
// 	count := 0
// 	maxHeight := 0
// 	for left < right {
// 		for left < right && height[left] <= maxHeight {
// 			left++
// 		}

// 		for left < right && height[right] <= maxHeight {
// 			right--
// 		}

// 		min := int(min(float64(height[left]), float64(height[right])))
// 		for i := left + 1; i < right; i++ {
// 			if height[i] < min {
// 				if height[i] < maxHeight {
// 					count += min - maxHeight
// 				} else {
// 					count += min - height[i]
// 				}
// 			}
// 		}

// 		maxHeight = int(max(float64(maxHeight), float64(min)))
// 	}

// 	return count
// }

// Brute force
// func trap(height []int) int {
// 	n := len(height)
// 	count := 0

// 	for i := 0; i < n; i++ {
// 		j := i
// 		leftMax := 0
// 		rightMax := 0

// 		for j >= 0 {
// 			leftMax = int(max(float64(leftMax), float64(height[j])))
// 			j--
// 		}

// 		j = i
// 		for j < n {
// 			rightMax = int(max(float64(rightMax), float64(height[j])))
// 			j++
// 		}

// 		count += int(min(float64(leftMax), float64(rightMax))) - height[i]
// 	}

// 	return count
// }

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 1, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
	fmt.Println(trap([]int{6, 4, 2, 0, 3, 2, 0, 3, 1, 4, 5, 3, 2, 7, 5, 3, 0, 1, 2, 1, 3, 4, 6, 8, 1, 3}))
	fmt.Println(trap([]int{2, 0, 2}))
}
