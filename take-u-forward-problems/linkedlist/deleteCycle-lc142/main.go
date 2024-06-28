package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/starting-point-of-loop-in-a-linked-list/
// Using tortoise and hare
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next

		if slow == fast {
			slow = head
			for slow != fast {
				slow, fast = slow.Next, fast.Next
			}

			return slow
		}
	}

	return nil
}

// // Using hashing
// func detectCycle(head *ListNode) *ListNode {
// 	list := map[*ListNode]int{}

// 	for head != nil {
// 		_, ok := list[head]
// 		if ok {
// 			return head
// 		}

// 		list[head] = 1
// 		head = head.Next
// 	}

// 	return nil
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

	list1.Next.Next.Next.Next.Next = list1.Next.Next

	fmt.Println(detectCycle(list1))
}
