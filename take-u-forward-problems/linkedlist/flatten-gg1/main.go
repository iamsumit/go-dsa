package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/flattening-a-linked-list/
// Merge sort approach
func flattenNode(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}

	return merge(node, flattenNode(node.Next))
}

func merge(l1, l2 *ListNode) *ListNode {
	sortedNode := &ListNode{}
	current := sortedNode

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Bottom = l1
			current = l1
			l1 = l1.Bottom
		} else {
			current.Bottom = l2
			current = l2
			l2 = l2.Bottom
		}

		current.Next = nil
	}

	if l1 != nil {
		current.Bottom = l1
	} else {
		current.Bottom = l2
	}

	if sortedNode.Bottom != nil {
		sortedNode.Bottom.Next = nil
	}

	return sortedNode.Bottom
}

// Brute force
// func flattenNode(node *ListNode) *ListNode {
// 	list := []int{}
// 	for l := node; l != nil; l = l.Next {
// 		list = append(list, l.Val)
// 		for b := l.Bottom; b != nil; b = b.Bottom {
// 			list = append(list, b.Val)
// 		}
// 	}

// 	slices.Sort(list)

// 	sortedNode := &ListNode{Val: 0}
// 	tmp := sortedNode
// 	for i := 0; i < len(list); i++ {
// 		tmp.Bottom = &ListNode{
// 			Val: list[i],
// 		}

// 		tmp = tmp.Bottom
// 	}

// 	return sortedNode.Bottom
// }

type ListNode struct {
	Val    int
	Next   *ListNode
	Bottom *ListNode
}

func main() {
	list := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Bottom: &ListNode{
				Val: 10,
			},
			Next: &ListNode{
				Val: 1,
				Bottom: &ListNode{
					Val: 7,
					Bottom: &ListNode{
						Val: 11,
						Bottom: &ListNode{
							Val: 12,
						},
					},
				},
				Next: &ListNode{
					Val: 4,
					Bottom: &ListNode{
						Val: 9,
					},
					Next: &ListNode{
						Val: 5,
						Bottom: &ListNode{
							Val: 6,
							Bottom: &ListNode{
								Val: 8,
							},
						},
					},
				},
			},
		},
	}

	list2 := flattenNode(list)

	fmt.Println("---------------")
	for l := list2; l != nil; l = l.Bottom {
		fmt.Println(l.Val)
	}
	fmt.Println("---------------")
}
