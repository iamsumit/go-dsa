package main

import (
	"fmt"
)

// https://leetcode.com/problems/top-k-frequent-elements/

// Approach 3 - non heap
func topKFrequent(nums []int, k int) []int {
	hashMap := map[int]int{}
	for _, val := range nums {
		hashMap[val]++
	}

	concurrency := make([][]int, len(nums)+1)
	for digit, freq := range hashMap {
		concurrency[freq] = append(concurrency[freq], digit)
	}

	output := []int{}
	for i := len(concurrency) - 1; i >= 0; i-- {
		if k <= 0 {
			return output
		}

		if len(concurrency[i]) == 0 {
			continue
		}

		output = append(output, concurrency[i]...)
		k -= len(concurrency[i])
	}

	return output
}

// type Item struct {
// 	Val   int
// 	Count int
// }

// type MaxHeap []Item

// func (h MaxHeap) Len() int           { return len(h) }
// func (h MaxHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
// func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *MaxHeap) Push(x interface{}) {
// 	*h = append(*h, x.(Item))
// }

// func (h *MaxHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// Approach 2
// func topKFrequent(nums []int, k int) []int {
// 	hashMap := map[int]int{}
// 	for _, val := range nums {
// 		hashMap[val]++
// 	}

// 	m := &MaxHeap{}
// 	heap.Init(m)
// 	for val, count := range hashMap {
// 		heap.Push(m, Item{
// 			Val:   val,
// 			Count: count,
// 		})
// 	}

// 	output := []int{}
// 	for k > 0 {
// 		item := heap.Pop(m).(Item)
// 		output = append(output, item.Val)
// 		k--
// 	}

// 	return output
// }

// Approach 1
// func topKFrequent(nums []int, k int) []int {
// 	hashMap := map[int]int{}
// 	m := &MaxHeap{}
// 	heap.Init(m)
// 	for _, val := range nums {
// 		hashMap[val]++

// 		heap.Push(m, Item{
// 			Val:   val,
// 			Count: hashMap[val],
// 		})
// 	}

// 	output := []int{}
// 	for k > 0 {
// 		item := heap.Pop(m).(Item)
// 		if _, ok := hashMap[item.Val]; ok {
// 			output = append(output, item.Val)
// 			k--
// 			delete(hashMap, item.Val)
// 		}
// 	}

// 	return output
// }

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 3))
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 1))
	fmt.Println(topKFrequent([]int{1}, 1))
	fmt.Println(topKFrequent([]int{1, 2}, 2))
	fmt.Println(topKFrequent([]int{3, 0, 1, 0}, 1))
}
