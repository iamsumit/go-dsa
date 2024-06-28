package main

import "fmt"

// https://takeuforward.org/data-structure/delete-given-node-in-a-linked-list-o1-approach/
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
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
	}

	node := head.Next

	deleteNode(node)
	for h := head; h != nil; h = h.Next {
		fmt.Println(h.Val)
	}
}
