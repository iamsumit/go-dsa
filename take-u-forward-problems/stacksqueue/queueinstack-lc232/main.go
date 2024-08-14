package main

import "fmt"

// https://takeuforward.org/data-structure/implement-queue-using-stack/
type MyStack []int

func (m *MyStack) Push(data int) {
	*m = append(*m, data)
}

// LIFO
func (m *MyStack) Pop() int {
	vals := *m
	n := len(vals)
	if len(vals) == 0 {
		return -1
	}

	x := vals[n-1]
	*m = vals[0 : n-1]
	return x
}

func (m *MyStack) Top() int {
	vals := *m
	return vals[len(vals)-1]
}

func (m *MyStack) Empty() bool {
	return len(*m) == 0
}

type MyQueue struct {
	s  MyStack
	s2 MyStack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (m *MyQueue) Push(x int) {
	// length := len(m.s)
	for !m.s.Empty() {
		m.s2.Push(m.s.Pop())
	}

	m.s2.Push(x)

	for !m.s2.Empty() {
		m.s.Push(m.s2.Pop())
	}

	fmt.Println("m.s", m.s, "m.s2", m.s2)
}

func (m *MyQueue) Pop() int {
	return m.s.Pop()
}

func (m *MyQueue) Peek() int {
	return m.s.Top()
}

func (m *MyQueue) Empty() bool {
	return m.s.Empty()
}

func main() {
	myQueue := Constructor()
	myQueue.Push(1)
	myQueue.Push(2)
	fmt.Println(myQueue.Peek())
	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Empty())
}
