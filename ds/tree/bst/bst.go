package bst

import "fmt"

type BinarySearchTreeNode struct {
	Data  int
	Left  *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

func NewBinarySearchTreeNode(data int) *BinarySearchTreeNode {
	return &BinarySearchTreeNode{Data: data}
}

func (node *BinarySearchTreeNode) Insert(data int) {
	if data < node.Data {
		if node.Left == nil {
			node.Left = NewBinarySearchTreeNode(data)
		} else {
			node.Left.Insert(data)
		}
	} else {
		if node.Right == nil {
			node.Right = NewBinarySearchTreeNode(data)
		} else {
			node.Right.Insert(data)
		}
	}
}

func (node *BinarySearchTreeNode) GetLeft() *BinarySearchTreeNode {
	return node.Left
}

func (node *BinarySearchTreeNode) GetRight() *BinarySearchTreeNode {
	return node.Right
}

func (node *BinarySearchTreeNode) GetData() int {
	return node.Data
}

func (node *BinarySearchTreeNode) Search(data int) bool {
	if node == nil {
		return false
	}

	if data < node.Data {
		return node.Left.Search(data)
	} else if data > node.Data {
		return node.Right.Search(data)
	}

	return true
}

func (node *BinarySearchTreeNode) InOrderTraversal() {
	if node == nil {
		return
	}

	node.Left.InOrderTraversal()
	fmt.Printf("%d, ", node.Data)
	node.Right.InOrderTraversal()
}

func (node *BinarySearchTreeNode) PreOrderTraversal() {
	if node == nil {
		return
	}

	fmt.Printf("%d, ", node.Data)
	node.Left.PreOrderTraversal()
	node.Right.PreOrderTraversal()
}

func (node *BinarySearchTreeNode) PostOrderTraversal() {
	if node == nil {
		return
	}

	node.Left.PostOrderTraversal()
	node.Right.PostOrderTraversal()
	fmt.Printf("%d, ", node.Data)
}
