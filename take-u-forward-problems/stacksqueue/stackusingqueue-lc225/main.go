package main

import "fmt"

// https://takeuforward.org/data-structure/implement-stack-using-single-queue/
type MyQueue struct {
	val []int
}

func (m *MyQueue) Push(x int) {
	m.val = append(m.val, x)
}

func (m *MyQueue) Pop() int {
	if len(m.val) == 0 {
		return -1
	}

	x := m.val[0]
	m.val = m.val[1:]
	return x
}

func (m *MyQueue) Top() int {
	if len(m.val) == 0 {
		return -1
	}

	return m.val[0]
}

func (m *MyQueue) Empty() bool {
	return len(m.val) == 0
}

type MyStack struct {
	q MyQueue
}

func Constructor() MyStack {
	return MyStack{}
}

func (m *MyStack) Push(x int) {
	length := len(m.q.val)
	m.q.Push(x)
	for i := 0; i < length; i++ {
		m.q.Push(m.q.Pop())
	}
}

func (m *MyStack) Pop() int {
	return m.q.Pop()
}

func (m *MyStack) Top() int {
	return m.q.Top()
}

func (m *MyStack) Empty() bool {
	return m.q.Empty()
}

func main() {
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.Top())
	myStack.Pop()
	fmt.Println(myStack.Empty())
}
