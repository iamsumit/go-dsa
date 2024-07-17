package main

import "fmt"

// https://takeuforward.org/data-structure/search-single-element-in-a-sorted-array/

// Approach 2
func singleNonDuplicate(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	} else if nums[0] != nums[1] {
		return nums[0]
	} else if nums[length-1] != nums[length-2] {
		return nums[length-1]
	}

	low, high := 0, length-2

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] != nums[mid+1] && nums[mid] != nums[mid-1] {
			return nums[mid]
		}

		if (mid%2 == 1 && nums[mid] == nums[mid-1]) || (mid%2 == 0 && nums[mid] == nums[mid+1]) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// Approach 1
// func singleNonDuplicate(nums []int) int {
// 	low, high := 0, len(nums)-1

// 	for low < high {
// 		mid := (low + high) / 2
// 		if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
// 			return nums[mid]
// 		}

// 		if mid%2 == 0 {
// 			if nums[mid] != nums[mid-1] {
// 				low = mid
// 			} else {
// 				high = mid
// 			}
// 		} else {
// 			if nums[mid] == nums[mid-1] {
// 				low = mid + 1
// 			} else {
// 				high = mid - 1
// 			}
// 		}
// 	}

// 	return nums[low]
// }

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 2, 2, 3, 3}))             // 1
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 2, 3, 3, 4, 8, 8})) // 4
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8})) // 2
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 2, 3, 4, 4, 8, 8})) // 3
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))    // 10
	fmt.Println(singleNonDuplicate([]int{3, 3, 4, 7, 7, 11, 11}))     // 4
	fmt.Println(singleNonDuplicate([]int{1}))                         // 1
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 2, 3}))             // 3
	fmt.Println(singleNonDuplicate([]int{7, 7, 10, 11, 11, 12, 12}))  // 10
}
