package main

import (
	"fmt"
	"math"
)

// https://takeuforward.org/data-structure/median-of-two-sorted-arrays-of-different-sizes/

// approach 2
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	n := n1 + n2

	left := (n1 + n2 + 1) / 2
	low, high := 0, n1
	for low <= high {
		mid1 := (low + high) / 2
		mid2 := left - mid1

		l1, l2, r1, r2 := math.MinInt64, math.MinInt64, math.MaxInt64, math.MaxInt64
		if mid1 < n1 {
			r1 = nums1[mid1]
		}

		if mid2 < n2 {
			r2 = nums2[mid2]
		}

		if mid1-1 >= 0 {
			l1 = nums1[mid1-1]
		}

		if mid2-1 >= 0 {
			l2 = nums2[mid2-1]
		}

		if l1 <= r2 && l2 <= r1 {
			if n%2 == 1 {
				return float64(max(l1, l2))
			} else {
				return (float64(max(l1, l2)) + float64(min(r1, r2))) / 2
			}
		} else if l1 > r2 {
			high = mid1 - 1
		} else {
			low = mid1 + 1
		}
	}

	return 0
}

// approach 1
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	m, n := len(nums1), len(nums2)
// 	oddCount := ((m+n)%2 == 1)

// 	if m == 0 && n == 0 {
// 		return 0
// 	} else if m != 0 && n == 0 {
// 		return median(nums1)
// 	} else if m == 0 && n != 0 {
// 		return median(nums2)
// 	}

// 	low, high := min(nums1[0], nums2[0]), max(nums1[m-1], nums2[n-1])
// 	req := (m + n) / 2
// 	for low <= high {
// 		mid := (low + high) / 2
// 		count := minCount(nums1, mid) + minCount(nums2, mid)
// 		if count <= req {
// 			low = mid + 1
// 		} else {
// 			high = mid - 1
// 		}
// 	}

// 	if oddCount {
// 		return float64(low)
// 	}

// 	elem1 := low

// 	low, high = min(nums1[0], nums2[0]), max(nums1[m-1], nums2[n-1])

// 	req2 := req - 1
// 	for low <= high {
// 		mid := (low + high) / 2
// 		count := minCount(nums1, mid) + minCount(nums2, mid)
// 		if count <= req2 {
// 			low = mid + 1
// 		} else {
// 			high = mid - 1
// 		}
// 	}

// 	elem2 := low

// 	return float64(elem1+elem2) / 2
// }

// func median(arr []int) float64 {
// 	length := len(arr)
// 	if length%2 == 1 {
// 		return float64(arr[length/2])
// 	}

// 	return float64((arr[length/2] + arr[(length/2)-1])) / 2
// }

// func minCount(arr []int, midX int) int {
// 	left, right := 0, len(arr)-1
// 	ans := len(arr)
// 	for left <= right {
// 		mid := (left + right) / 2
// 		if arr[mid] > midX {
// 			ans = mid
// 			right = mid - 1
// 		} else {
// 			left = mid + 1
// 		}
// 	}

// 	return ans
// }

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 4}, []int{2, 3}))
	fmt.Println(findMedianSortedArrays([]int{1, 4, 5}, []int{2, 3}))
	fmt.Println(findMedianSortedArrays([]int{1, 4, 5}, []int{2, 3, 6}))
	fmt.Println(findMedianSortedArrays([]int{2, 2, 4, 4}, []int{2, 2, 4, 4}))
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0}))
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{2, 3}))
}
