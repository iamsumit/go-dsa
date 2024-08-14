package main

import "fmt"

// https://takeuforward.org/data-structure/implement-queue-using-array/
type MyQueue []int

func (m *MyQueue) Push(data int) {
	*m = append(*m, data)
}

// FIFO
func (m *MyQueue) Pop() {
	vals := *m
	if len(vals) == 0 {
		return
	}

	*m = vals[1:]
}

func (m *MyQueue) Top() int {
	vals := *m
	return vals[0]
}

func main() {
	myQueue := &MyQueue{}
	myQueue.Push(2)
	myQueue.Push(3)
	fmt.Println("myQueue", myQueue)
	fmt.Println("myQueue.Top()", myQueue.Top())
	fmt.Println("myQueue.Pop()")
	myQueue.Pop()
	fmt.Println("myQueue.Push(4)")
	myQueue.Push(4)
	fmt.Println("myQueue", myQueue)
	fmt.Println("myQueue.Top()", myQueue.Top())
	myQueue.Pop()
	fmt.Println("myQueue.Pop()")
	fmt.Println("myQueue", myQueue)
	fmt.Println("myQueue.Top()", myQueue.Top())
}
