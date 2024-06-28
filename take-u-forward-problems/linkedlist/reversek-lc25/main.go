package main

import "fmt"

// https://takeuforward.org/data-structure/reverse-linked-list-in-groups-of-size-k/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}

	revNode := &ListNode{}
	tmp := revNode

	for head != nil {
		tmp.Next, head = getKthReverseNode(head, k)

		for tmp.Next != nil {
			tmp = tmp.Next
		}
	}

	return revNode.Next
}

func getKthReverseNode(node *ListNode, k int) (*ListNode, *ListNode) {
	partNode := &ListNode{}

	curr := node
	tmp := partNode
	for i := 0; i < k; i++ {
		if i+1 <= k && curr == nil {
			return node, nil
		}

		next := curr.Next
		curr.Next = nil
		tmp.Next = curr
		tmp = tmp.Next

		curr = next
	}

	return reverseNode(partNode.Next), curr
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
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: 8,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val: 10,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	head := reverseKGroup(list, 4)
	for h := head; h != nil; h = h.Next {
		fmt.Println(h.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
