package main

import "fmt"

// https://takeuforward.org/data-structure/reverse-a-linked-list/
// Reverse using recursion
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)

	next := head.Next
	next.Next = head
	head.Next = nil

	return newHead
}

// Reverse using loop
// func reverseList(head *ListNode) *ListNode {
// 	var prev *ListNode

// 	curr := head

// 	for curr != nil {
// 		next := curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 	}

// 	return prev
// }

func main() {
	head := reverseList(&ListNode{
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

type ListNode struct {
	Val  int
	Next *ListNode
}
