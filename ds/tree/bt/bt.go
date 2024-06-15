package bt

import "fmt"

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func NewBinaryTreeNode(data int) *BinaryTreeNode {
	return &BinaryTreeNode{Data: data}
}

func (node *BinaryTreeNode) SetLeft(left *BinaryTreeNode) {
	node.Left = left
}

func (node *BinaryTreeNode) SetRight(right *BinaryTreeNode) {
	node.Right = right
}

func (node *BinaryTreeNode) GetLeft() *BinaryTreeNode {
	return node.Left
}

func (node *BinaryTreeNode) GetRight() *BinaryTreeNode {
	return node.Right
}

func (node *BinaryTreeNode) GetData() int {
	return node.Data
}

func (node *BinaryTreeNode) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}

func (node *BinaryTreeNode) IsFull() bool {
	return node.Left != nil && node.Right != nil
}

func (node *BinaryTreeNode) Delete(data int, parent *BinaryTreeNode) bool {
	if node == nil {
		return false
	}

	if node.Data == data {
		if node.IsLeaf() {
			if parent != nil {
				if parent.Left == node {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
			}

			return true
		} else if node.Left == nil {
			node.Data = node.Right.Data
			node.Left = node.Right.Left
			node.Right = node.Right.Right
			return true
		} else if node.Right == nil {
			node.Data = node.Left.Data
			node.Right = node.Left.Right
			node.Left = node.Left.Left
			return true
		} else {
			temp := node.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			node.Data = temp.Data
			return node.Right.Delete(temp.Data, node)
		}
	}

	if node.Left.Delete(data, node) {
		return true
	}

	if node.Right.Delete(data, node) {
		return true
	}

	return false
}

func (node *BinaryTreeNode) Search(data int) bool {
	if node == nil {
		return false
	}

	if node.Data == data {
		return true
	}

	if node.Left.Search(data) {
		return true
	}

	if node.Right.Search(data) {
		return true
	}

	return false
}

func (node *BinaryTreeNode) InOrderTraversal() {
	if node == nil {
		return
	}

	node.Left.InOrderTraversal()
	fmt.Printf("%d, ", node.Data)
	node.Right.InOrderTraversal()
}

func (node *BinaryTreeNode) PreOrderTraversal() {
	if node == nil {
		return
	}

	fmt.Printf("%d, ", node.Data)
	node.Left.PreOrderTraversal()
	node.Right.PreOrderTraversal()
}

func (node *BinaryTreeNode) PostOrderTraversal() {
	if node == nil {
		return
	}

	node.Left.PostOrderTraversal()
	node.Right.PostOrderTraversal()
	fmt.Printf("%d, ", node.Data)
}
