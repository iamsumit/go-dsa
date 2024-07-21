package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/kth-largest-smallest-element-in-an-array/
// Heap sort approach
type Heap struct {
	Val []int
}

func (h *Heap) Insert(val int) {
	h.Val = append(h.Val, val)
	heapifyUp(h, len(h.Val)-1)
}

func (h *Heap) Pop() {
	h.Val[0], h.Val[len(h.Val)-1] = h.Val[len(h.Val)-1], h.Val[0]
	h.Val = h.Val[:len(h.Val)-1]
	heapifyDown(h, 0)
}

func heapifyUp(h *Heap, index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.Val[index] > h.Val[parent] {
			h.Val[index], h.Val[parent] = h.Val[parent], h.Val[index]
			index = parent
		} else {
			break
		}
	}
}

func heapifyDown(h *Heap, index int) {
	for {
		right_child := (index + 1) * 2
		left_child := right_child - 1

		largest := index

		if left_child < len(h.Val) && h.Val[left_child] > h.Val[largest] {
			largest = left_child
		}

		if right_child < len(h.Val) && h.Val[right_child] > h.Val[largest] {
			largest = right_child
		}

		if largest != index {
			h.Val[index], h.Val[largest] = h.Val[largest], h.Val[index]
			index = largest
		} else {
			break
		}
	}
}

func findKthLargest(nums []int, k int) int {
	var heap Heap
	for _, val := range nums {
		heap.Insert(val)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop()
	}

	return heap.Val[0]
}

// Quick sort
// func findKthLargest(nums []int, k int) int {
// 	quickSort(nums, 0, len(nums)-1)
// 	return nums[len(nums)-k]
// }

// func quickSort(arr []int, low, high int) {
// 	if low < high {
// 		pivot := partition(arr, low, high)

// 		quickSort(arr, low, pivot-1)
// 		quickSort(arr, pivot+1, high)
// 	}
// }

// func partition(arr []int, low, high int) int {
// 	pivot := arr[high]
// 	i := low - 1

// 	for j := low; j < high; j++ {
// 		if arr[j] < pivot {
// 			i++
// 			arr[i], arr[j] = arr[j], arr[i]
// 		}

// 	}

// 	arr[i+1], arr[high] = arr[high], arr[i+1]

// 	return i + 1
// }

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
