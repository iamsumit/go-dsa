package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/find-median-from-data-stream/

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	lowerHalf *MaxHeap
	upperHalf *MinHeap
}

func Constructor() MedianFinder {
	lowerHalf := &MaxHeap{}
	upperHalf := &MinHeap{}
	heap.Init(lowerHalf)
	heap.Init(upperHalf)
	return MedianFinder{lowerHalf, upperHalf}
}

func (mf *MedianFinder) AddNum(num int) {
	heap.Push(mf.lowerHalf, num)
	heap.Push(mf.upperHalf, heap.Pop(mf.lowerHalf))

	if mf.lowerHalf.Len() < mf.upperHalf.Len() {
		heap.Push(mf.lowerHalf, heap.Pop(mf.upperHalf))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.lowerHalf.Len() > mf.upperHalf.Len() {
		return float64((*mf.lowerHalf)[0])
	}
	return float64((*mf.lowerHalf)[0]+(*mf.upperHalf)[0]) / 2.0
}

func main() {
	medianFinder := Constructor()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)
	fmt.Println(medianFinder.FindMedian())
	medianFinder.AddNum(3)
	fmt.Println(medianFinder.FindMedian())
	fmt.Println("-----------------")

	// medianFinder2 := Constructor()
	// medianFinder2.AddNum(0)
	// medianFinder2.AddNum(0)
	// fmt.Println(medianFinder2.FindMedian())
	// fmt.Println("-----------------")

	medianFinder3 := Constructor()
	medianFinder3.AddNum(6)
	fmt.Println(medianFinder3.FindMedian())
	medianFinder3.AddNum(10)
	fmt.Println(medianFinder3.FindMedian())
	medianFinder3.AddNum(2)
	fmt.Println(medianFinder3.FindMedian())
	medianFinder3.AddNum(6)
	fmt.Println(medianFinder3.FindMedian())
	medianFinder3.AddNum(5)
	fmt.Println(medianFinder3.FindMedian())
	medianFinder3.AddNum(0)
	fmt.Println(medianFinder3.FindMedian())
	medianFinder3.AddNum(6)
	fmt.Println(medianFinder3.FindMedian())
	medianFinder3.AddNum(3)
	fmt.Println(medianFinder3.FindMedian())
	medianFinder3.AddNum(1)
	fmt.Println(medianFinder3.FindMedian())
	medianFinder3.AddNum(0)
	fmt.Println(medianFinder3.FindMedian())
	medianFinder3.AddNum(0)
	fmt.Println(medianFinder3.FindMedian())
	fmt.Println("-----------------")

	medianFinder4 := Constructor()
	medianFinder4.AddNum(1)
	fmt.Println(medianFinder4.FindMedian())
	medianFinder4.AddNum(2)
	fmt.Println(medianFinder4.FindMedian())
	medianFinder4.AddNum(3)
	fmt.Println(medianFinder4.FindMedian())
	medianFinder4.AddNum(4)
	fmt.Println(medianFinder4.FindMedian())
	medianFinder4.AddNum(5)
	fmt.Println(medianFinder4.FindMedian())
	medianFinder4.AddNum(6)
	fmt.Println(medianFinder4.FindMedian())
	medianFinder4.AddNum(7)
	fmt.Println(medianFinder4.FindMedian())
	medianFinder4.AddNum(8)
	fmt.Println(medianFinder4.FindMedian())
	medianFinder4.AddNum(9)
	fmt.Println(medianFinder4.FindMedian())
	medianFinder4.AddNum(10)
	fmt.Println(medianFinder4.FindMedian())
}
