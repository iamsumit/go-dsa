package main

// Optimized solution
func sortColors(nums []int) {
	left, mid, right := 0, 0, len(nums)-1

	for mid <= right {
		if nums[mid] == 0 {
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[mid], nums[right] = nums[right], nums[mid]
			right--
		}
	}
}

// Using heap sort
// func sortColors(nums []int) {
// 	n := len(nums)
// 	for i := n/2 - 1; i >= 0; i-- {
// 		heapify(nums, n, i)
// 	}

// 	for i := n - 1; i > 0; i-- {
// 		nums[0], nums[i] = nums[i], nums[0]
// 		heapify(nums, i, 0)
// 	}
// }

// func heapify(nums []int, n, i int) {
// 	largest := i
// 	l := 2*i + 1
// 	r := 2*i + 2

// 	if l < n && nums[l] > nums[largest] {
// 		largest = l
// 	}

// 	if r < n && nums[r] > nums[largest] {
// 		largest = r
// 	}

// 	if largest != i {
// 		nums[i], nums[largest] = nums[largest], nums[i]
// 		heapify(nums, n, largest)
// 	}
// }

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0}) // [0 0 1 1 2 2]
	sortColors([]int{2, 0, 1})          // [0 1 2]
	sortColors([]int{0})                // [0]
	sortColors([]int{1})                // [1]
}
