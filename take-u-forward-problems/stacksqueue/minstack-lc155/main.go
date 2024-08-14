package main

import (
	"fmt"
	"math"
)

// https://takeuforward.org/data-structure/implement-min-stack-o2n-and-on-space-complexity/

// Approach 3
type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   math.MaxInt64,
	}
}

func (m *MinStack) Push(value int) {
	if len(m.stack) == 0 {
		m.min = value
		m.stack = append(m.stack, value)
	} else {
		if value < m.min {
			m.stack = append(m.stack, 2*value-m.min)
			m.min = value
		} else {
			m.stack = append(m.stack, value)
		}
	}
}

func (m *MinStack) Pop() {
	if len(m.stack) == 0 {
		return
	}

	val := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	if val < m.min {
		m.min = 2*m.min - val
	}
}

func (m *MinStack) Top() int {
	val := m.stack[len(m.stack)-1]
	if val < m.min {
		return int(m.min)
	}

	return int(val)
}

func (m *MinStack) GetMin() int {
	return int(m.min)
}

// // Approach 2
// type Node struct {
// 	val    int
// 	minVal int
// }
// type MinStack struct {
// 	data []Node
// }

// func Constructor() MinStack {
// 	return MinStack{
// 		data: []Node{},
// 	}
// }

// func (m *MinStack) Push(val int) {
// 	node := Node{
// 		val: val,
// 	}

// 	if len(m.data) == 0 {
// 		node.minVal = val
// 	} else {
// 		node.minVal = min(m.GetMin(), val)
// 	}

// 	m.data = append(m.data, node)
// }

// func (m *MinStack) Pop() {
// 	if len(m.data) == 0 {
// 		return
// 	}

// 	m.data = m.data[:len(m.data)-1]
// }

// func (m *MinStack) Top() int {
// 	return m.data[len(m.data)-1].val
// }

// func (m *MinStack) GetMin() int {
// 	return m.data[len(m.data)-1].minVal
// }

// type MinStack struct {
// 	data    []int
// 	minElem int
// }

// func Constructor() MinStack {
// 	return MinStack{
// 		data:    []int{},
// 		minElem: math.MaxInt,
// 	}
// }

// func (m *MinStack) Push(val int) {
// 	m.data = append(m.data, val)
// 	m.minElem = min(m.minElem, val)
// }

// func (m *MinStack) Pop() {
// 	if len(m.data) == 0 {
// 		return
// 	}

// 	val := m.data[len(m.data)-1]
// 	m.data = m.data[:len(m.data)-1]
// 	if m.minElem != val {
// 		return
// 	}

// 	m.minElem = math.MaxInt
// 	for _, val := range m.data {
// 		m.minElem = min(m.minElem, val)
// 	}
// }

// func (m *MinStack) Top() int {
// 	return m.data[len(m.data)-1]
// }

// func (m *MinStack) GetMin() int {
// 	return m.minElem
// }

// Brute force approach 1
// type MinStack struct {
// 	data    []int
// 	minElem int
// }

// func Constructor() MinStack {
// 	return MinStack{
// 		data:    []int{},
// 		minElem: math.MaxInt,
// 	}
// }

// func (m *MinStack) Push(val int) {
// 	m.data = append(m.data, val)
// 	m.minElem = min(m.minElem, val)
// }

// func (m *MinStack) Pop() {
// 	if len(m.data) == 0 {
// 		return
// 	}

// 	val := m.data[len(m.data)-1]
// 	m.data = m.data[:len(m.data)-1]
// 	if len(m.data) == 0 {
// 		m.minElem = math.MaxInt
// 	} else {
// 		if m.minElem == val {
// 			m.minElem = m.data[0]
// 			for _, v := range m.data {
// 				m.minElem = min(m.minElem, v)
// 			}
// 		}
// 	}
// }

// func (m *MinStack) Top() int {
// 	return m.data[len(m.data)-1]
// }

// func (m *MinStack) GetMin() int {
// 	return m.minElem
// }

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	obj := Constructor()
	// obj.Pop()
	// obj.Push(-2)
	// obj.Push(0)
	// obj.Push(-3)
	// fmt.Println(obj.GetMin())
	// obj.Pop()
	// fmt.Println(obj.Top())
	// fmt.Println(obj.GetMin())

	// obj.Push(-10)
	// fmt.Println("obj.Push(-10)")
	// obj.Push(14)
	// fmt.Println("obj.Push(14)")
	// fmt.Println("obj.GetMin()", obj.GetMin())
	// fmt.Println("obj.GetMin()", obj.GetMin())
	// obj.Push(-20)
	// fmt.Println("obj.Push(-20)")
	// fmt.Println("obj.GetMin()", obj.GetMin())
	// fmt.Println("obj.GetMin()", obj.GetMin())
	// fmt.Println("obj.Top()", obj.Top())
	// fmt.Println("obj.GetMin()", obj.GetMin())
	// obj.Pop()
	// fmt.Println("obj.Pop()")
	// fmt.Println("obj.GetMin()", obj.GetMin())
	// obj.Push(10)
	// fmt.Println("obj.Push(10)")
	// obj.Push(-7)
	// fmt.Println("obj.Push(-7)")
	// fmt.Println("obj.GetMin()", obj.GetMin())
	// obj.Push(-7)
	// fmt.Println("obj.Push(-7)")
	// obj.Pop()
	// fmt.Println("obj.Pop()")
	// fmt.Println("obj.Top()", obj.Top())
	// fmt.Println("obj.GetMin()", obj.GetMin())
	// obj.Pop()
	// fmt.Println("obj.Pop()")

	obj.Push(2)
	fmt.Println("obj.Push(2)")
	obj.Push(0)
	fmt.Println("obj.Push(0)")
	obj.Push(3)
	fmt.Println("obj.Push(3)")
	obj.Push(0)
	fmt.Println("obj.Push(0)")
	fmt.Println("obj.GetMin()", obj.GetMin())
	obj.Pop()
	fmt.Println("obj.Pop()")
	fmt.Println("obj.GetMin()", obj.GetMin())
	obj.Pop()
	fmt.Println("obj.Pop()")
	fmt.Println("obj.GetMin()", obj.GetMin())
	obj.Pop()
	fmt.Println("obj.Pop()")
	fmt.Println("obj.GetMin()", obj.GetMin())
}
