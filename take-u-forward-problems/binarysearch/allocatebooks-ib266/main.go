package main

import (
	"fmt"
	"slices"
)

// https://takeuforward.org/data-structure/allocate-minimum-number-of-pages/

// binary search
func books(books []int, students int) int {
	if students > len(books) {
		return -1
	}

	low := slices.Max(books)
	high := sumArray(books)

	for low <= high {
		pages := (low + high) / 2
		count := countStudents(books, pages)

		if count > students {
			low = pages + 1
		} else {
			high = pages - 1
		}
	}

	return low
}

// Brute force
// func books(books []int, students int) int {
// 	if students > len(books) {
// 		return -1
// 	}

// 	low := slices.Max(books)
// 	high := sumArray(books)

// 	for pages := low; pages <= high; pages++ {
// 		if countStudents(books, pages) == students {
// 			return pages
// 		}
// 	}

// 	return low
// }

func countStudents(arr []int, pages int) int {
	n := len(arr)
	students := 1
	pagesStudent := 0
	for i := 0; i < n; i++ {
		if pagesStudent+arr[i] <= pages {
			pagesStudent += arr[i]
		} else {
			students += 1
			pagesStudent = arr[i]
		}
	}

	return students
}

func sumArray(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}

	return sum
}

func main() {
	fmt.Println(books([]int{12, 34, 67, 90}, 2))
	fmt.Println(books([]int{5, 17, 100, 11}, 4))
	fmt.Println(books([]int{25, 46, 28, 49, 24}, 4))
	fmt.Println(books([]int{73, 58, 30, 72, 44, 78, 23, 9}, 5))
	fmt.Println(books([]int{100, 100, 100}, 2))
}
