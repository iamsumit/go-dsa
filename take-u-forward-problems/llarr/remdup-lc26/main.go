package main

import "fmt"

// Optimal approach 2
func removeDuplicates(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[count] != nums[i] {
			count++
			nums[count] = nums[i]
		}
	}

	return count + 1
}

// Optimal approach
// func removeDuplicates(nums []int) int {
// 	count := 1
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i-1] != nums[i] {
// 			nums[count] = nums[i]
// 			count++
// 		}
// 	}

// 	return count
// }

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
