package nry

import "fmt"

type NaryTreenode struct {
	Data     int
	Children []*NaryTreenode
}

func NewNaryTreeNode(data int) *NaryTreenode {
	return &NaryTreenode{Data: data}
}

func (node *NaryTreenode) AddChild(child *NaryTreenode) {
	node.Children = append(node.Children, child)
}

func (node *NaryTreenode) GetChildren() []*NaryTreenode {
	return node.Children
}

func (node *NaryTreenode) GetData() int {
	return node.Data
}

func (node *NaryTreenode) IsLeaf() bool {
	return len(node.Children) == 0
}

func (node *NaryTreenode) Search(data int) bool {
	if node == nil {
		return false
	}

	if node.Data == data {
		return true
	}

	for _, child := range node.Children {
		if child.Search(data) {
			return true
		}
	}

	return false
}

func (node *NaryTreenode) GetChildCount() int {
	return len(node.Children)
}

func (node *NaryTreenode) GetChild(index int) *NaryTreenode {
	if index < 0 || index >= len(node.Children) {
		return nil
	}

	return node.Children[index]
}

func (node *NaryTreenode) InOrderTraversal() {
	if node == nil {
		return
	}

	if len(node.Children) > 0 {
		node.Children[0].InOrderTraversal()
	}

	fmt.Printf("%d, ", node.Data)

	for i := 1; i < len(node.Children); i++ {
		node.Children[i].InOrderTraversal()
	}
}

func (node *NaryTreenode) PreOrderTraversal() {
	if node == nil {
		return
	}

	fmt.Printf("%d, ", node.Data)

	for _, child := range node.Children {
		child.PreOrderTraversal()
	}
}

func (node *NaryTreenode) PostOrderTraversal() {
	if node == nil {
		return
	}

	for _, child := range node.Children {
		child.PostOrderTraversal()
	}

	fmt.Printf("%d, ", node.Data)
}
