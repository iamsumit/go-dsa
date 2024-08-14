package main

import (
	"fmt"
)

type MyStack []int

func (m *MyStack) Push(x int) {
	*m = append(*m, x)
}

func (m *MyStack) Pop() int {
	vals := *m
	x := vals[len(vals)-1]
	*m = vals[:len(vals)-1]
	return x
}

func (m *MyStack) Top() int {
	vals := *m
	return vals[len(vals)-1]
}

func (m *MyStack) IsEmpty() bool {
	return len(*m) == 0
}

func (m *MyStack) Len() int {
	return len(*m)
}

// https://www.naukri.com/code360/problems/sort-a-stack_985275?topList=striver-sde-sheet-problems&utm_source=striver&utm_medium=website
func sortStack(stack MyStack) {
	var sortedInsert func(stack *MyStack, val int)
	sortedInsert = func(stack *MyStack, val int) {
		if stack.Len() == 0 || val > stack.Top() {
			stack.Push(val)
			return
		}

		tmp := stack.Pop()
		sortedInsert(stack, val)
		stack.Push(tmp)
	}

	if stack.Len() != 0 {
		tmp := stack.Pop()
		sortStack(stack)
		sortedInsert(&stack, tmp)
	}
}

func main() {
	stack := []int{5, -2, 9, -7, 3}
	sortStack(MyStack(stack))
	fmt.Println(stack)
}
