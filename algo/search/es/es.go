package es

import "github.com/iamsumit/go-dsa/algo/search/bs"

func Search(nums []int, target int) int {
	if len(nums) > 0 && nums[0] == target {
		return 0
	}

	startAt := 0
	endAt := len(nums)

	for i := 1; i < len(nums); i = i * 2 {
		if nums[i] == target {
			return i
		}

		if nums[i] > target {
			startAt = i / 2
			endAt = i
		}

		if i > len(nums) {
			startAt = i / 2
		}
	}

	return bs.Search(nums[startAt:endAt], target)
}
