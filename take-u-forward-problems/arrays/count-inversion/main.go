package main

import "fmt"

// With merge sort
func sort(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(arr []int, l, r int) int {
	count := 0
	if l < r {
		m := l + (r-l)/2
		count += mergeSort(arr, l, m)
		count += mergeSort(arr, m+1, r)

		count += merge(arr, l, m, r)
	}

	return count
}

func merge(arr []int, l, m, r int) int {
	count := 0
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
			count += (m + 1 - l - i)
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

	return count
}

// Brute force
// func sort(nums []int) int {
// 	var count int

// 	for i := 0; i < len(nums); i++ {
// 		for k := i; k < len(nums); k++ {
// 			if nums[k] < nums[i] {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }

func main() {
	fmt.Println(sort([]int{1, 2, 3, 4, 5}))
	fmt.Println(sort([]int{5, 4, 3, 2, 1}))
	fmt.Println(sort([]int{5, 3, 2, 1, 4}))
}
