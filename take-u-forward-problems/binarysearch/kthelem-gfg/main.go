package main

import (
	"fmt"
	"math"
)

// https://takeuforward.org/data-structure/k-th-element-of-two-sorted-arrays/
// approach 2
func kthElem(arr1 []int, arr2 []int, k int) int {
	m, n := len(arr1), len(arr2)
	if m > n {
		return kthElem(arr2, arr1, k)
	}

	low, high := max(0, k-n), min(k, m)
	for low <= high {
		mid1 := (low + high) / 2
		mid2 := k - mid1

		l1, l2, r1, r2 := math.MinInt, math.MinInt, math.MaxInt, math.MaxInt
		if mid1 < m {
			r1 = arr1[mid1]
		}

		if mid2 < n {
			r2 = arr2[mid2]
		}

		if mid1-1 >= 0 {
			l1 = arr1[mid1-1]
		}

		if mid2-1 >= 0 {
			l2 = arr2[mid2-1]
		}

		if l1 <= r2 && l2 <= r1 {
			return max(l1, l2)
		} else if l1 > r2 {
			high = mid1 - 1
		} else {
			low = mid1 + 1
		}
	}

	return 0
}

// // First approach
// func kthElem(arr1 []int, arr2 []int, k int) int {
// 	m, n := len(arr1), len(arr2)
// 	switch {
// 	case m == 0 && n == 0:
// 		return -1
// 	case m == 0 && n != 0:
// 		return arr2[k-1]
// 	case m != 0 && n == 0:
// 		return arr1[k-1]
// 	}

// 	low, high := min(arr1[0], arr2[0]), max(arr1[m-1], arr2[n-1])
// 	for low <= high {
// 		mid := (low + high) / 2
// 		count := minCount(arr1, mid) + minCount(arr2, mid)
// 		if count >= k {
// 			high = mid - 1
// 		} else {
// 			low = mid + 1
// 		}
// 	}

// 	return low
// }

// func minCount(nums []int, midX int) int {
// 	ans := len(nums)
// 	low, high := 0, len(nums)-1
// 	for low <= high {
// 		mid := (low + high) / 2
// 		if nums[mid] > midX {
// 			ans = mid
// 			high = mid - 1
// 		} else {
// 			low = mid + 1
// 		}
// 	}

// 	return ans
// }

func main() {
	fmt.Println(kthElem([]int{2, 3, 6, 7, 9}, []int{1, 4, 8, 10}, 5))
	fmt.Println(kthElem([]int{100, 112, 256, 349, 770}, []int{72, 86, 113, 119, 265, 445, 892}, 7))
	fmt.Println(kthElem([]int{}, []int{1, 4, 8, 10}, 2))
}
