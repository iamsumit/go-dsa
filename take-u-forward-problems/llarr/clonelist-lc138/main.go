package main

import (
	"fmt"
)

// Without hashmap
func copyRandomList(head *Node) *Node {
	tmp := head
	for tmp != nil {
		nextElem := tmp.Next
		newNode := &Node{
			Val:  tmp.Val,
			Next: nextElem,
		}

		tmp.Next = newNode
		tmp = nextElem
	}

	tmp = head
	for tmp != nil {
		newNode := tmp.Next
		if tmp.Random != nil {
			newNode.Random = tmp.Random.Next
		} else {
			newNode.Random = nil
		}

		tmp = tmp.Next.Next
	}

	newNode := &Node{}
	dummy := newNode

	tmp = head
	for tmp != nil {
		dummy.Next = tmp.Next
		dummy = dummy.Next
		tmp.Next = tmp.Next.Next
		tmp = tmp.Next
	}

	return newNode.Next
}

// Using hashmap
// func copyRandomList(head *Node) *Node {
// 	newNode := &Node{}
// 	tmp := newNode

// 	hashMap := map[*Node]*Node{}

// 	curr := head
// 	for curr != nil {
// 		tmp.Next = &Node{
// 			Val: curr.Val,
// 		}

// 		hashMap[curr] = tmp.Next

// 		curr = curr.Next
// 		tmp = tmp.Next
// 	}

// 	curr = head
// 	tmp = newNode
// 	for curr != nil {
// 		r, ok := hashMap[curr.Random]
// 		if ok {
// 			tmp.Next.Random = r
// 		}

// 		tmp = tmp.Next
// 		curr = curr.Next
// 	}

// 	return newNode.Next
// }

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func main() {
	list := &Node{
		Val: 7,
		Next: &Node{
			Val: 13,
			Next: &Node{
				Val: 11,
				Next: &Node{
					Val: 10,
					Next: &Node{
						Val: 1,
					},
				},
			},
		},
	}

	list.Next.Random = list
	list.Next.Next.Random = list.Next.Next.Next.Next
	list.Next.Next.Next.Random = list.Next.Next
	list.Next.Next.Next.Next.Random = list

	tmp := list
	head := copyRandomList(list)
	for h := head; h != nil; h = h.Next {
		var rand, hrand int
		if tmp.Random != nil {
			rand = tmp.Random.Val
		}
		if h.Random != nil {
			hrand = h.Random.Val
		}

		fmt.Println(h == tmp, h.Random == tmp.Random, "vals:", h.Val, tmp.Val, "random:", hrand, rand)

		tmp = tmp.Next
	}
}
