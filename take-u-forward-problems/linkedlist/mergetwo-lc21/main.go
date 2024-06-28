package main

import "fmt"

// https://takeuforward.org/data-structure/merge-two-sorted-linked-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	sortedNode := &ListNode{}
	current := sortedNode

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return sortedNode.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	new := mergeTwoLists(list1, list2)

	fmt.Println("printint list...")
	for i := new; i != nil; i = i.Next {
		fmt.Println("val", i.Val)
	}

}
