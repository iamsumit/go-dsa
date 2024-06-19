package is

func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		pos := left + ((right-left)/(nums[right]-nums[left]))*(target-nums[left])

		if pos > right || pos < left {
			return -1
		}

		if nums[pos] == target {
			return pos
		} else if nums[pos] < target {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}

	return -1
}
