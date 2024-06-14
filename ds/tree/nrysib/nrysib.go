package nrysib

import "fmt"

type TreeNode struct {
	Data        int
	FirstChild  *TreeNode
	NextSibling *TreeNode
}

func NewTreeNode(data int) *TreeNode {
	return &TreeNode{Data: data}
}

func (node *TreeNode) AddChild(child *TreeNode) {
	if node.FirstChild == nil {
		node.FirstChild = child
	} else {
		node.FirstChild.AddSibling(child)
	}
}

func (node *TreeNode) AddSibling(sibling *TreeNode) {
	if node.NextSibling == nil {
		node.NextSibling = sibling
	} else {
		node.NextSibling.AddSibling(sibling)
	}
}

func (node *TreeNode) GetSibling() *TreeNode {
	return node.NextSibling
}

func (node *TreeNode) GetChildren() *TreeNode {
	return node.FirstChild
}

func (node *TreeNode) GetData() int {
	return node.Data
}

func (node *TreeNode) GetChildCount() int {
	count := 0
	child := node.FirstChild
	for child != nil {
		count++
		child = child.NextSibling
	}

	return count
}

func (node *TreeNode) GetChild(index int) *TreeNode {
	if index < 0 {
		return nil
	}

	child := node.FirstChild
	for i := 0; i < index && child != nil; i++ {
		child = child.NextSibling
	}

	return child
}

func (node *TreeNode) Search(data int) bool {
	if node == nil {
		return false
	}

	if node.Data == data {
		return true
	}

	if node.FirstChild != nil {
		if node.FirstChild.Search(data) {
			return true
		}
	}

	if node.NextSibling != nil {
		if node.NextSibling.Search(data) {
			return true
		}
	}

	return false
}

func (node *TreeNode) PreOrderTraversal() {
	if node == nil {
		return
	}

	fmt.Printf("%d, ", node.Data)

	if node.FirstChild != nil {
		node.FirstChild.PreOrderTraversal()
	}

	if node.NextSibling != nil {
		node.NextSibling.PreOrderTraversal()
	}
}

func (node *TreeNode) PostOrderTraversal() {
	if node == nil {
		return
	}

	if node.FirstChild != nil {
		node.FirstChild.PostOrderTraversal()
	}

	fmt.Printf("%d, ", node.Data)

	if node.NextSibling != nil {
		node.NextSibling.PostOrderTraversal()
	}

}
