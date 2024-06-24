package main

import "fmt"

// Merge sort - optimal
func reversePairs(nums []int) int {
	return divideAndConcur(nums, 0, len(nums)-1)
}

func divideAndConcur(nums []int, left, right int) int {
	count := 0
	if left < right {
		mid := left + (right-left)/2
		count += divideAndConcur(nums, left, mid)
		count += divideAndConcur(nums, mid+1, right)
		count += countPair(nums, left, mid, right)
		merge(nums, left, mid, right)
	}

	return count
}

func countPair(nums []int, left, mid, right int) int {
	count := 0
	high := mid + 1

	for i := left; i <= mid; i++ {
		for high <= right && nums[i] > nums[high]*2 {
			high++
		}

		count += (high - (mid + 1))
	}

	return count
}

func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	L := make([]int, n1)
	R := make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}

	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	i, j, k := 0, 0, l

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}

}

// Brute force
// func reversePairs(nums []int) int {
// 	count := 0

// 	left := 0
// 	right := len(nums) - 1
// 	for left < right {
// 		if nums[left] > nums[right]*2 {
// 			count++
// 		}

// 		left++

// 		if left == right {
// 			left = 0
// 			right--
// 		}
// 	}

// 	return count
// }

func main() {
	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
	fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
	fmt.Println(reversePairs([]int{2, 4, 3, 1, 5}))
	fmt.Println(reversePairs([]int{2, 4, 3, 1, 0}))
	fmt.Println(reversePairs([]int{-5, -5}))
}
