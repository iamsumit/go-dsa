package main

import (
	"testing"

	"github.com/iamsumit/go-dsa/ds/tree/bst"
	"github.com/iamsumit/go-dsa/ds/tree/bt"
	"github.com/iamsumit/go-dsa/ds/tree/tt"
)

func BenchmarkBinaryTreeCreationWithTwoLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		btVar := bt.NewBinaryTreeNode(1)
		btVar.SetLeft(bt.NewBinaryTreeNode(2))
		btVar.SetRight(bt.NewBinaryTreeNode(3))
	}
}

func BenchmarkBinaryTreeUpdateIn31Node(b *testing.B) {
	btVar := getBinaryTree()
	for i := 0; i < b.N; i++ {
		btVar.SetLeft(bt.NewBinaryTreeNode(i))
	}
}

func BenchmarkBinaryTreeSearchIn31Node(b *testing.B) {
	btVar := getBinaryTree()
	for i := 0; i < b.N; i++ {
		btVar.Search(i)
	}
}

func BenchmarkBinarySearchTreeInsertWithHugeNode(b *testing.B) {
	bstVar := bst.NewBinarySearchTreeNode(b.N / 2)

	for i := 0; i < b.N; i++ {
		bstVar.Insert(i)
	}
}

func BenchmarkBinarySearchTreeSearchIn31Node(b *testing.B) {
	bstVar := bst.NewBinarySearchTreeNode(15)

	for i := 0; i < 31; i++ {
		bstVar.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		bstVar.Search(i)
	}
}

func BenchmarkBinarySearchTreeSearchIn10000Node(b *testing.B) {
	bstVar := bst.NewBinarySearchTreeNode(5000)

	for i := 0; i < 10000; i++ {
		bstVar.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		bstVar.Search(i)
	}
}

func BenchmarkTernaryTreeInsertWithTwoLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ttVar := tt.NewTernaryTreeNode(i)
		ttVar.SetLeft(tt.NewTernaryTreeNode(i + 1))
		ttVar.SetMid(tt.NewTernaryTreeNode(i + 2))
		ttVar.SetRight(tt.NewTernaryTreeNode(i + 3))
	}
}

func BenchmarkTernaryTreeUpdateWith40Nodes(b *testing.B) {
	ttVar := getTernaryTree()
	for i := 0; i < b.N; i++ {
		ttVar.SetLeft(tt.NewTernaryTreeNode(i))
	}
}

func BenchmarkTernaryTreeSearchWith40Nodes(b *testing.B) {
	ttVar := getTernaryTree()
	for i := 0; i < b.N; i++ {
		ttVar.Search(i)
	}
}

func getBinaryTree() *bt.BinaryTreeNode {
	btVar := bt.NewBinaryTreeNode(1)

	btVar.SetLeft(bt.NewBinaryTreeNode(2))
	btVar.SetRight(bt.NewBinaryTreeNode(3))

	btVar.GetLeft().SetLeft(bt.NewBinaryTreeNode(4))
	btVar.GetLeft().SetRight(bt.NewBinaryTreeNode(5))
	btVar.GetRight().SetLeft(bt.NewBinaryTreeNode(6))
	btVar.GetRight().SetRight(bt.NewBinaryTreeNode(7))

	btVar.GetLeft().GetLeft().SetLeft(bt.NewBinaryTreeNode(8))
	btVar.GetLeft().GetLeft().SetRight(bt.NewBinaryTreeNode(9))
	btVar.GetLeft().GetRight().SetLeft(bt.NewBinaryTreeNode(10))
	btVar.GetLeft().GetRight().SetRight(bt.NewBinaryTreeNode(11))

	btVar.GetRight().GetLeft().SetLeft(bt.NewBinaryTreeNode(12))
	btVar.GetRight().GetLeft().SetRight(bt.NewBinaryTreeNode(13))
	btVar.GetRight().GetRight().SetLeft(bt.NewBinaryTreeNode(14))
	btVar.GetRight().GetRight().SetRight(bt.NewBinaryTreeNode(15))

	btVar.GetLeft().GetLeft().GetLeft().SetLeft(bt.NewBinaryTreeNode(16))
	btVar.GetLeft().GetLeft().GetLeft().SetRight(bt.NewBinaryTreeNode(17))
	btVar.GetLeft().GetLeft().GetRight().SetLeft(bt.NewBinaryTreeNode(18))
	btVar.GetLeft().GetLeft().GetRight().SetRight(bt.NewBinaryTreeNode(19))

	btVar.GetLeft().GetRight().GetLeft().SetLeft(bt.NewBinaryTreeNode(20))
	btVar.GetLeft().GetRight().GetLeft().SetRight(bt.NewBinaryTreeNode(21))
	btVar.GetLeft().GetRight().GetRight().SetLeft(bt.NewBinaryTreeNode(22))
	btVar.GetLeft().GetRight().GetRight().SetRight(bt.NewBinaryTreeNode(23))

	btVar.GetRight().GetLeft().GetLeft().SetLeft(bt.NewBinaryTreeNode(24))
	btVar.GetRight().GetLeft().GetLeft().SetRight(bt.NewBinaryTreeNode(25))
	btVar.GetRight().GetLeft().GetRight().SetLeft(bt.NewBinaryTreeNode(26))
	btVar.GetRight().GetLeft().GetRight().SetRight(bt.NewBinaryTreeNode(27))

	btVar.GetRight().GetRight().GetLeft().SetLeft(bt.NewBinaryTreeNode(28))
	btVar.GetRight().GetRight().GetLeft().SetRight(bt.NewBinaryTreeNode(29))
	btVar.GetRight().GetRight().GetRight().SetLeft(bt.NewBinaryTreeNode(30))
	btVar.GetRight().GetRight().GetRight().SetRight(bt.NewBinaryTreeNode(31))
	return btVar
}

func getTernaryTree() *tt.TernaryTreeNode {
	ttVar := tt.NewTernaryTreeNode(1)
	ttVar.SetLeft(tt.NewTernaryTreeNode(2))
	ttVar.SetMid(tt.NewTernaryTreeNode(3))
	ttVar.SetRight(tt.NewTernaryTreeNode(4))

	ttVar.GetLeft().SetLeft(tt.NewTernaryTreeNode(5))
	ttVar.GetLeft().SetMid(tt.NewTernaryTreeNode(6))
	ttVar.GetLeft().SetRight(tt.NewTernaryTreeNode(7))

	ttVar.GetMid().SetLeft(tt.NewTernaryTreeNode(8))
	ttVar.GetMid().SetMid(tt.NewTernaryTreeNode(9))
	ttVar.GetMid().SetRight(tt.NewTernaryTreeNode(10))

	ttVar.GetRight().SetLeft(tt.NewTernaryTreeNode(11))
	ttVar.GetRight().SetMid(tt.NewTernaryTreeNode(12))
	ttVar.GetRight().SetRight(tt.NewTernaryTreeNode(13))

	ttVar.GetLeft().GetLeft().SetLeft(tt.NewTernaryTreeNode(14))
	ttVar.GetLeft().GetLeft().SetMid(tt.NewTernaryTreeNode(15))
	ttVar.GetLeft().GetLeft().SetRight(tt.NewTernaryTreeNode(16))

	ttVar.GetLeft().GetMid().SetLeft(tt.NewTernaryTreeNode(17))
	ttVar.GetLeft().GetMid().SetMid(tt.NewTernaryTreeNode(18))
	ttVar.GetLeft().GetMid().SetRight(tt.NewTernaryTreeNode(19))

	ttVar.GetLeft().GetRight().SetLeft(tt.NewTernaryTreeNode(20))
	ttVar.GetLeft().GetRight().SetMid(tt.NewTernaryTreeNode(21))
	ttVar.GetLeft().GetRight().SetRight(tt.NewTernaryTreeNode(22))

	ttVar.GetMid().GetLeft().SetLeft(tt.NewTernaryTreeNode(23))
	ttVar.GetMid().GetLeft().SetMid(tt.NewTernaryTreeNode(24))
	ttVar.GetMid().GetLeft().SetRight(tt.NewTernaryTreeNode(25))

	ttVar.GetMid().GetMid().SetLeft(tt.NewTernaryTreeNode(26))
	ttVar.GetMid().GetMid().SetMid(tt.NewTernaryTreeNode(27))
	ttVar.GetMid().GetMid().SetRight(tt.NewTernaryTreeNode(28))

	ttVar.GetMid().GetRight().SetLeft(tt.NewTernaryTreeNode(29))
	ttVar.GetMid().GetRight().SetMid(tt.NewTernaryTreeNode(30))
	ttVar.GetMid().GetRight().SetRight(tt.NewTernaryTreeNode(31))

	ttVar.GetRight().GetLeft().SetLeft(tt.NewTernaryTreeNode(32))
	ttVar.GetRight().GetLeft().SetMid(tt.NewTernaryTreeNode(33))
	ttVar.GetRight().GetLeft().SetRight(tt.NewTernaryTreeNode(34))

	ttVar.GetRight().GetMid().SetLeft(tt.NewTernaryTreeNode(35))
	ttVar.GetRight().GetMid().SetMid(tt.NewTernaryTreeNode(36))
	ttVar.GetRight().GetMid().SetRight(tt.NewTernaryTreeNode(37))

	ttVar.GetRight().GetRight().SetLeft(tt.NewTernaryTreeNode(38))
	ttVar.GetRight().GetRight().SetMid(tt.NewTernaryTreeNode(39))
	ttVar.GetRight().GetRight().SetRight(tt.NewTernaryTreeNode(40))

	return ttVar
}
