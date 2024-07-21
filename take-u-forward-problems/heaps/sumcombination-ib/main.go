package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://www.interviewbit.com/problems/maximum-sum-combinations/
type Pair struct {
	sum int
	i   int
	j   int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].sum > h[j].sum }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// appraoch 2
func solve(arr1 []int, arr2 []int, length int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)

	n := len(arr1)
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	visited := make(map[[2]int]bool)

	heap.Push(maxHeap, Pair{arr1[n-1] + arr2[n-1], n - 1, n - 1})
	visited[[2]int{n - 1, n - 1}] = true

	results := []int{}

	for len(results) < length {
		top := heap.Pop(maxHeap).(Pair)
		results = append(results, top.sum)

		i, j := top.i, top.j

		if i-1 >= 0 && !visited[[2]int{i - 1, j}] {
			heap.Push(maxHeap, Pair{arr1[i-1] + arr2[j], i - 1, j})
			visited[[2]int{i - 1, j}] = true
		}

		if j-1 >= 0 && !visited[[2]int{i, j - 1}] {
			heap.Push(maxHeap, Pair{arr1[i] + arr2[j-1], i, j - 1})
			visited[[2]int{i, j - 1}] = true
		}
	}

	return results
}

// Approach 1
// func solve(arr1 []int, arr2 []int) []int {
// 	sum := []int{}
// 	for _, val := range arr1 {
// 		for _, val2 := range arr2 {
// 			sum = append(sum, val+val2)
// 		}
// 	}

// 	var heap Heap
// 	for _, val := range sum {
// 		heap.Insert(val)
// 	}

// 	output := []int{}
// 	for i := 0; i < len(arr1); i++ {
// 		output = append(output, heap.Pop())
// 	}

// 	return output
// }

func main() {
	fmt.Println(solve([]int{3, 2}, []int{1, 4}, 2))
	fmt.Println(solve([]int{1, 4, 2, 3}, []int{2, 5, 1, 6}, 4))
	fmt.Println(solve([]int{59, 63, 65, 6, 46, 82, 28, 62, 92, 96, 43, 28, 37, 92, 5, 3, 54, 93, 83}, []int{59, 63, 65, 6, 46, 82, 28, 62, 92, 96, 43, 28, 37, 92, 5, 3, 54, 93, 83}, 10))
	fmt.Println(solve([]int{58, 38, 60}, []int{58, 38, 60}, 2))
}
