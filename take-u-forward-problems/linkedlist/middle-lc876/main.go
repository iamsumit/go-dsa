package main

import "fmt"

// https://takeuforward.org/data-structure/find-middle-element-in-a-linked-list/
// Tortoise and Hare algorithm
func middleNode(head *ListNode) *ListNode {
	tortoise, hare := head, head

	for hare != nil && hare.Next != nil && tortoise.Next != nil {
		hare, tortoise = hare.Next.Next, tortoise.Next
	}

	return tortoise
}

// Brute force with arr
// func middleNode(head *ListNode) *ListNode {
// 	arr := []int{}

// 	for h := head; h != nil; h = h.Next {
// 		arr = append(arr, h.Val)
// 	}

// 	mid := len(arr)/2 + 1

// 	for i := 0; i < mid-1; i++ {
// 		head = head.Next
// 	}

// 	return head
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := middleNode(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	})
	for h := head; h != nil; h = h.Next {
		fmt.Println(h.Val)
	}
}
