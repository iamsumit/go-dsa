package main

import "fmt"

// https://takeuforward.org/data-structure/add-two-numbers-represented-as-linked-lists/
func addTwoNumbers(list1 *ListNode, list2 *ListNode) *ListNode {
	sumNode := &ListNode{}
	current := sumNode
	carry := 0

	for list1 != nil || list2 != nil {
		sum := 0

		if list1 != nil {
			sum += list1.Val
			list1 = list1.Next
		}

		if list2 != nil {
			sum += list2.Val
			list2 = list2.Next
		}

		sum += carry
		carry = sum / 10

		current.Next = &ListNode{
			Val: sum % 10,
		}

		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{
			Val: 1,
		}
	}

	return sumNode.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	list2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	new := addTwoNumbers(list1, list2)

	fmt.Println("printint list...")
	for i := new; i != nil; i = i.Next {
		fmt.Println("val", i.Val)
	}

}
