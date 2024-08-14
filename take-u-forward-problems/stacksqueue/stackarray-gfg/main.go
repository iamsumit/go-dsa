package main

import "fmt"

// https://takeuforward.org/data-structure/implement-stack-using-array/
type MyStack []int

func (m *MyStack) Push(data int) {
	*m = append(*m, data)
}

// LIFO
func (m *MyStack) Pop() {
	vals := *m
	if len(vals) == 0 {
		return
	}

	*m = vals[0 : len(vals)-1]
}

func (m *MyStack) Top() int {
	vals := *m
	return vals[len(vals)-1]
}

func main() {
	myStack := &MyStack{}
	myStack.Push(2)
	myStack.Push(3)
	fmt.Println("myStack", myStack)
	fmt.Println("myStack.Top()", myStack.Top())
	fmt.Println("myStack.Pop()")
	myStack.Pop()
	fmt.Println("myStack.Push(4)")
	myStack.Push(4)
	fmt.Println("myStack.Top()", myStack.Top())
	myStack.Pop()
	fmt.Println("myStack", myStack)
	fmt.Println("myStack.Top()", myStack.Top())
}
