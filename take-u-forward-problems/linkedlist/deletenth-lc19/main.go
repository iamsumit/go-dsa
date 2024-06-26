package main

import "fmt"

// with loop
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}

	h := head
	for i := 0; i < n; i++ {
		if i+1 <= n && h == nil {
			return head
		}

		h = h.Next
	}

	if h == nil {
		return head.Next
	}

	curr := head
	for h.Next != nil {
		curr = curr.Next
		h = h.Next
	}

	curr.Next = curr.Next.Next

	return head
}

// With array
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	var list []*ListNode

// 	curr := head
// 	for curr != nil {
// 		list = append(list, curr)
// 		curr = curr.Next
// 	}

// 	index := len(list) - n
// 	if index == 0 {
// 		return head.Next
// 	} else if index+1 >= len(list) {
// 		list[index-1].Next = nil
// 	} else {
// 		list[index-1].Next = list[index+1]
// 	}

// 	return head
// }

// With recursion
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	i := delete(head, 0, n)
// 	if i == n {
// 		return head.Next
// 	}

// 	return head
// }

// func delete(head *ListNode, i, n int) int {
// 	if head == nil || head.Next == nil {
// 		return 1
// 	}

// 	i = delete(head.Next, i, n)

// 	if i == n {
// 		head.Next = head.Next.Next
// 	}

// 	return i + 1
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := removeNthFromEnd(&ListNode{
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
	}, 1)
	for h := head; h != nil; h = h.Next {
		fmt.Println(h.Val)
	}
}
