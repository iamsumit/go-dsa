package main

import "fmt"

// https://takeuforward.org/data-structure/find-the-duplicate-in-an-array-of-n1-integers/
func findDuplicate(nums []int) int {
	prev := nums[0]
	next := nums[0]

	for {
		prev = nums[prev]       // 3 2 4
		next = nums[nums[next]] // 4 4 4

		if prev == next {
			break
		}
	}

	next = nums[0]
	for prev != next {
		prev = nums[prev]
		next = nums[next]
	}

	return prev
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicate([]int{3, 3, 3, 3, 3}))
	fmt.Println(findDuplicate([]int{1, 2, 3, 2}))
}
