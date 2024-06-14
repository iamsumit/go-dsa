package tt

import "fmt"

type TernaryTreeNode struct {
	Data  int
	Left  *TernaryTreeNode
	Mid   *TernaryTreeNode
	Right *TernaryTreeNode
}

func NewTernaryTreeNode(data int) *TernaryTreeNode {
	return &TernaryTreeNode{Data: data}
}

func (t *TernaryTreeNode) GetData() int {
	return t.Data
}

func (t *TernaryTreeNode) SetLeft(node *TernaryTreeNode) {
	t.Left = node
}

func (t *TernaryTreeNode) GetLeft() *TernaryTreeNode {
	return t.Left
}

func (t *TernaryTreeNode) SetMid(node *TernaryTreeNode) {
	t.Mid = node
}

func (t *TernaryTreeNode) GetMid() *TernaryTreeNode {
	return t.Mid
}

func (t *TernaryTreeNode) SetRight(node *TernaryTreeNode) {
	t.Right = node
}

func (t *TernaryTreeNode) GetRight() *TernaryTreeNode {
	return t.Right
}

func (t *TernaryTreeNode) IsLeaf() bool {
	return t.Left == nil && t.Mid == nil && t.Right == nil
}

func (t *TernaryTreeNode) IsFull() bool {
	return t.Left != nil && t.Mid != nil && t.Right != nil
}

func (t *TernaryTreeNode) Search(data int) bool {
	if t == nil {
		return false
	}

	if t.Data == data {
		return true
	}

	if t.Left.Search(data) {
		return true
	}

	if t.Mid.Search(data) {
		return true
	}

	if t.Right.Search(data) {
		return true
	}

	return false
}

func (t *TernaryTreeNode) InOrderTraversal() {
	if t == nil {
		return
	}

	t.Left.InOrderTraversal()
	fmt.Printf("%d, ", t.Data)
	t.Mid.InOrderTraversal()
	t.Right.InOrderTraversal()
}

func (t *TernaryTreeNode) PreOrderTraversal() {
	if t == nil {
		return
	}

	fmt.Printf("%d, ", t.Data)
	t.Left.PreOrderTraversal()
	t.Mid.PreOrderTraversal()
	t.Right.PreOrderTraversal()
}

func (t *TernaryTreeNode) PostOrderTraversal() {
	if t == nil {
		return
	}

	t.Left.PostOrderTraversal()
	t.Mid.PostOrderTraversal()
	t.Right.PostOrderTraversal()
	fmt.Printf("%d, ", t.Data)
}
