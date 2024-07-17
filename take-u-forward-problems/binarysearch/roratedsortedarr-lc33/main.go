package main

import "fmt"

// https://takeuforward.org/data-structure/search-element-in-a-rotated-sorted-array/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}

		if nums[left] <= nums[mid] {
			if target < nums[left] || target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target < nums[mid] || target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), "expeced", 4)
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 5), "expected", 1)
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 7), "expected", 3)
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 4), "expected", 0)
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 2), "expected", 6)
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3), "expected", -1)
	fmt.Println(search([]int{5, 6, 7, 0, 1, 2, 4}, 5), "expected", 0)
	fmt.Println(search([]int{5, 6, 7, 0, 1, 2, 4}, 6), "expected", 1)
	fmt.Println(search([]int{5, 6, 7, 0, 1, 2, 4}, 7), "expected", 2)
	fmt.Println(search([]int{5, 6, 7, 0, 1, 2, 4}, 0), "expected", 3)
	fmt.Println(search([]int{5, 6, 7, 0, 1, 2, 4}, 1), "expected", 4)
	fmt.Println(search([]int{5, 6, 7, 0, 1, 2, 4}, 2), "expected", 5)
	fmt.Println(search([]int{5, 6, 7, 0, 1, 2, 4}, 4), "expected", 6)
	fmt.Println(search([]int{1}, 0), "expected", -1)
	fmt.Println(search([]int{1}, 1), "expected", 0)
}
