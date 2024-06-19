package js

import "math"

func Search(nums []int, target int) int {
	step := int(math.Sqrt(float64(len(nums))))

	startAt := 0
	endAt := len(nums)

	for i := 0; i < len(nums); i += step {
		if nums[i] == target {
			return i
		} else if nums[i] > target && i != 0 {
			startAt = i - step
			endAt = i
			break
		}

		if i+step >= len(nums) {
			startAt = i
			break
		}
	}

	for j := startAt; j < endAt; j++ {
		if nums[j] == target {
			return j
		}
	}

	return -1
}
