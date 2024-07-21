package main

import (
	"container/heap"
	"fmt"
)

// https://www.naukri.com/code360/problems/merge-k-sorted-arrays_975379

type Item struct {
	Val       int
	ArrIndex  int
	ElemIndex int
}

type MinHeap []*Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// Solution 2
func mergeKSortedArrays(karrs [][]int) []int {
	m := make(MinHeap, 0)
	heap.Init(&m)

	for i, karr := range karrs {
		heap.Push(&m, &Item{
			Val:       karr[0],
			ArrIndex:  i,
			ElemIndex: 0,
		})
	}

	sorted := []int{}
	for m.Len() > 0 {
		item := heap.Pop(&m).(*Item)
		sorted = append(sorted, item.Val)

		if item.ElemIndex+1 < len(karrs[item.ArrIndex]) {
			heap.Push(&m, &Item{
				Val:       karrs[item.ArrIndex][item.ElemIndex+1],
				ArrIndex:  item.ArrIndex,
				ElemIndex: item.ElemIndex + 1,
			})
		}
	}

	return sorted
}

// Solution 1
// func mergeKSortedArrays(karrs [][]int) []int {
// 	var heap Heap
// 	for _, karr := range karrs {
// 		for _, val := range karr {
// 			heap.Insert(val)
// 		}
// 	}

// 	sortedArr := []int{}
// 	total := len(heap.Val)
// 	for i := 0; i < total; i++ {
// 		sortedArr = append(sortedArr, heap.Pop())
// 	}

// 	return sortedArr
// }

// type Heap struct {
// 	Val []int
// }

// func (h *Heap) Insert(val int) {
// 	h.Val = append(h.Val, val)
// 	heapifyUp(h, len(h.Val)-1)
// }

// func (h *Heap) Pop() int {
// 	if len(h.Val) == 0 {
// 		return 0
// 	}

// 	x := h.Val[0]
// 	h.Val[0], h.Val[len(h.Val)-1] = h.Val[len(h.Val)-1], h.Val[0]
// 	h.Val = h.Val[:len(h.Val)-1]
// 	heapifyDown(h, 0)
// 	return x
// }

// func heapifyUp(h *Heap, index int) {
// 	for index > 0 {
// 		parent := (index - 1) / 2
// 		if h.Val[index] < h.Val[parent] {
// 			h.Val[index], h.Val[parent] = h.Val[parent], h.Val[index]
// 			index = parent
// 		} else {
// 			break
// 		}
// 	}
// }

// func heapifyDown(h *Heap, index int) {
// 	for {
// 		right_child := (index + 1) * 2
// 		left_child := right_child - 1

// 		smallest := index

// 		if left_child < len(h.Val) && h.Val[left_child] < h.Val[smallest] {
// 			smallest = left_child
// 		}

// 		if right_child < len(h.Val) && h.Val[right_child] < h.Val[smallest] {
// 			smallest = right_child
// 		}

// 		if smallest != index {
// 			(h.Val)[index], h.Val[smallest] = h.Val[smallest], h.Val[index]
// 			index = smallest
// 		} else {
// 			break
// 		}
// 	}
// }

func main() {
	fmt.Println(mergeKSortedArrays([][]int{{3, 5, 9}, {1, 2, 3, 8}}))
	fmt.Println(mergeKSortedArrays([][]int{{1, 5, 9}, {45, 90}, {2, 6, 78, 100, 234}, {0}}))
}
