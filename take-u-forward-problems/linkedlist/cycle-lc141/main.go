package main

import (
	"fmt"
)

// Tortoise and Hare
func hasCycle(head *ListNode) bool {
	tortoise, hare := head, head

	for hare != nil && hare.Next != nil && tortoise.Next != nil {
		hare, tortoise = hare.Next.Next, tortoise.Next

		if hare == tortoise {
			return true
		}
	}

	return false
}

// Using hashing
// func hasCycle(head *ListNode) bool {
// 	list := map[*ListNode]int{}

// 	for head != nil {
// 		_, ok := list[head]
// 		if ok {
// 			return true
// 		}

// 		list[head] = 1
// 		head = head.Next
// 	}

// 	return false
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	list1.Next.Next.Next.Next.Next = list1

	fmt.Println(hasCycle(list1))
}
