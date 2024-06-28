package main

import "fmt"

// https://takeuforward.org/data-structure/check-if-given-linked-list-is-plaindrome/
// using reverse and match
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	for head != nil && fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	revSlow := reverseNode(slow)

	for revSlow != nil && head != nil {
		if revSlow.Val != head.Val {
			return false
		}

		revSlow = revSlow.Next
		head = head.Next
	}

	return true
}

func reverseNode(node *ListNode) *ListNode {
	var prev *ListNode

	curr := node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// Using array
// func isPalindrome(head *ListNode) bool {
// 	list := []int{}
// 	for head != nil {
// 		list = append(list, head.Val)
// 		head = head.Next
// 	}

// 	left, right := 0, len(list)-1

// 	for left <= right {
// 		if list[left] != list[right] {
// 			return false
// 		}

// 		left++
// 		right--
// 	}

// 	return true
// }

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}

	fmt.Println(isPalindrome(list))
}

type ListNode struct {
	Val  int
	Next *ListNode
}
