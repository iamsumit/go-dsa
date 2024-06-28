package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/merge-two-sorted-arrays-without-extra-space/
// Optimal approach
func merge(nums1 []int, m int, nums2 []int, n int) {
	last := len(nums1) - 1

	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[last] = nums1[m-1]
			m--
		} else {
			nums1[last] = nums2[n-1]
			n--
		}

		last--
	}

	for n > 0 {
		nums1[last] = nums2[n-1]
		n--
		last--
	}
}

// Inserion sort
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	for i := 0; i < len(nums2); i++ {
// 		nums1[m+i] = nums2[i]
// 	}

// 	for i := 1; i < len(nums1); i++ {
// 		key := nums1[i]

// 		j := i - 1
// 		for j >= 0 && nums1[j] > key {
// 			nums1[j+1] = nums1[j]
// 			j--
// 		}

// 		nums1[j+1] = key
// 	}
// }

// with merge sort
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	L := make([]int, m)
// 	R := make([]int, n)

// 	for i := range L {
// 		L[i] = nums1[i]
// 	}

// 	for i := range R {
// 		R[i] = nums2[i]
// 	}

// 	lk, rk, key := 0, 0, 0
// 	for lk < m && rk < n {
// 		if L[lk] <= R[rk] {
// 			nums1[key] = L[lk]
// 			lk++
// 		} else {
// 			nums1[key] = R[rk]
// 			rk++
// 		}

// 		key++
// 	}

// 	for lk < m {
// 		nums1[key] = L[lk]
// 		lk++
// 		key++
// 	}

// 	for rk < n {
// 		nums1[key] = R[rk]
// 		rk++
// 		key++
// 	}
// }

func main() {
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// nums2 := []int{2, 5, 6}

	// merge(nums1, 3, nums2, 3)
	// fmt.Println(nums1, nums2)

	// nums1 = []int{0}
	// nums2 = []int{1}

	// merge(nums1, 0, nums2, 1)
	// fmt.Println(nums1, nums2)

	// nums1 = []int{2, 0}
	// nums2 = []int{1}

	// merge(nums1, 1, nums2, 1)
	// fmt.Println(nums1, nums2)

	nums1 := []int{4, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}

	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1, nums2)
}
