package main

import (
	"fmt"

	"github.com/iamsumit/go-dsa/ds/stack/st"
)

func main() {
	fmt.Println("---------------------")
	fmt.Println("Stack implementation in Go")
	fmt.Println("---------------------")

	stack := st.Stack{}
	fmt.Println("Initiated, data", stack.Data())
	stack.Push(1)
	fmt.Println("Pushed 1, data", stack.Data())
	stack.Push(2)
	fmt.Println("Pushed 2, data", stack.Data())

	fmt.Println("---------------------")

	fmt.Println("Peek:", stack.Peek())
	fmt.Println("Size:", stack.Size())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("Peek:", stack.Peek())
	fmt.Println("Data:", stack.Data())

	fmt.Println("---------------------")

	fmt.Println("IsEmpty:", stack.IsEmpty())
	fmt.Println("Pop:", stack.Pop())
	fmt.Println("IsEmpty:", stack.IsEmpty())

	fmt.Println("---------------------")
}
