package ts

func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid1 := left + (right-left)/3
		mid2 := right - (right-left)/3

		if nums[mid1] == target {
			return mid1
		} else if nums[mid2] == target {
			return mid2
		} else if target < nums[mid1] {
			right = mid1 - 1
		} else if target > nums[mid2] {
			left = mid2 + 1
		} else {
			left = mid1 + 1
			right = mid2 - 1
		}

	}

	return -1
}
