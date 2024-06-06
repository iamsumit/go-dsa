package main

import (
	"fmt"

	"github.com/iamsumit/go-dsa/ds/linkedlist/ll"
)

func main() {
	list := ll.LinkedList{}

	fmt.Println("------------------")

	list.Insert(1)
	fmt.Println("Head after first insert:", list.Head.Data)
	fmt.Println("Tail after first insert:", list.Tail.Data)

	fmt.Println("------------------")

	list.Insert(2)
	fmt.Println("Head after second insert", list.Head.Data)
	fmt.Println("Tail after second insert", list.Tail.Data)

	fmt.Println("------------------")

	list.Insert(3)
	list.Insert(4)

	fmt.Println("Loop over linked list after adding two elements:")
	for n := list.Head; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}

	fmt.Println("------------------")

	list.Delete(3)
	fmt.Println("Loop over linked list after deleting third element:")
	for n := list.Head; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}

	fmt.Println("------------------")

	list.Insert(5)
	fmt.Println("Loop over linked list after adding new data:")
	for n := list.Head; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}

	fmt.Println("------------------")

	node := list.Search(2)
	fmt.Println("Search for 2:", node.Data)

	fmt.Println("------------------")

	list.Reverse()
	fmt.Println("Loop over linked list after reversing:")
	for n := list.Head; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}

	fmt.Println("------------------")
}
