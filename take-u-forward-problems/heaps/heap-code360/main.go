package main

import "fmt"

// https://www.naukri.com/code360/problems/min-heap_4691801?topList=striver-sde-sheet-problems&utm_source=striver&utm_medium=website
type Heap struct {
	Val []int
}

func (h *Heap) Insert(val int) {
	h.Val = append(h.Val, val)
	heapifyUp(h, len(h.Val)-1)
}

func (h *Heap) Delete() {
	if len(h.Val) == 0 {
		return
	}

	h.Val[0], h.Val[len(h.Val)-1] = h.Val[len(h.Val)-1], h.Val[0]
	h.Val = h.Val[:len(h.Val)-1]
	heapifyDown(h, 0)
}

func heapifyUp(h *Heap, index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.Val[index] < h.Val[parent] {
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

		smallest := index

		if left_child < len(h.Val) && h.Val[left_child] < h.Val[smallest] {
			smallest = left_child
		}

		if right_child < len(h.Val) && h.Val[right_child] < h.Val[smallest] {
			smallest = right_child
		}

		if smallest != index {
			(h.Val)[index], h.Val[smallest] = h.Val[smallest], h.Val[index]
			index = smallest
		} else {
			break
		}
	}
}

func minHeap(Q [][]int) []int {
	var heap Heap
	output := []int{}

	for _, val := range Q {
		if val[0] == 0 {
			heap.Insert(val[1])
		} else if val[0] != 1 {
			heap.Insert(val[0])
		} else {
			output = append(output, heap.Val[0])
			heap.Delete()
		}
	}

	return output
}

func main() {
	fmt.Println(minHeap([][]int{{0, 2}, {0, 1}, {1}}))
	fmt.Println(minHeap([][]int{{2}, {5}, {0, 5}, {1}, {0, 43}, {0, 15}, {0, 5}, {2}, {0, 4}, {1}}))
}
