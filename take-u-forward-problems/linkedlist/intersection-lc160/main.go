package main

import (
	"fmt"
)

// Using two pointer approach
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	list1, list2 := headA, headB

	for list1 != list2 {
		if list1 == nil {
			list1 = headB
		} else {
			list1 = list1.Next
		}

		if list2 == nil {
			list2 = headA
		} else {
			list2 = list2.Next
		}
	}

	return list1
}

// Using hashing
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	list1 := map[*ListNode]*ListNode{}

// 	for headA != nil {
// 		list1[headA] = headA
// 		headA = headA.Next
// 	}

// 	for headB != nil {
// 		_, ok := list1[headB]
// 		if ok {
// 			return headB
// 		}

// 		headB = headB.Next
// 	}

// 	return nil
// }

// Brute force approach
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	for list1 := headA; list1 != nil; list1 = list1.Next {
// 		for list2 := headB; list2 != nil; list2 = list2.Next {
// 			if list1 == list2 {
// 				return list2
// 			}
// 		}
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

	list2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: list1.Next.Next,
				},
			},
		},
	}

	new := getIntersectionNode(list1, list2)

	fmt.Println("printint list...")
	for i := new; i != nil; i = i.Next {
		fmt.Println("val", i.Val)
	}

}
