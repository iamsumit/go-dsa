package main

import (
	"fmt"
)

// Using counting without array
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	tmp := head
	length := 1

	for tmp.Next != nil {
		tmp = tmp.Next
		length++
	}

	tmp.Next = head

	end := length - (k % length)

	for end != 0 {
		tmp = tmp.Next
		end--
	}

	head = tmp.Next
	tmp.Next = nil

	return head
}

// Approach 1 using array
// func rotateRight(head *ListNode, k int) *ListNode {
// 	if k == 0 || head == nil || head.Next == nil {
// 		return head
// 	}

// 	list := []*ListNode{}

// 	for head != nil {
// 		list = append(list, head)
// 		head = head.Next
// 	}

// 	breakPoint := len(list) - (k % len(list))

// 	resList := &ListNode{}
// 	tmp := resList

// 	for _, l := range list[breakPoint:] {
// 		tmp.Next = l
// 		tmp = tmp.Next
// 	}

// 	for _, l := range list[0:breakPoint] {
// 		tmp.Next = l
// 		tmp = tmp.Next
// 	}

// 	tmp.Next = nil

// 	return resList.Next
// }

func main() {
	list := &ListNode{
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

	head := rotateRight(list, 2)
	for h := head; h != nil; h = h.Next {
		fmt.Println(h.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
