package main

import (
	"fmt"

	"github.com/iamsumit/go-dsa/ds/tree/bst"
	"github.com/iamsumit/go-dsa/ds/tree/bt"
	"github.com/iamsumit/go-dsa/ds/tree/nry"
	"github.com/iamsumit/go-dsa/ds/tree/nrysib"
	"github.com/iamsumit/go-dsa/ds/tree/tt"
)

func main() {
	fmt.Println("Implmeneing Binary Tree...")
	fmt.Println("------------------------")
	btVar := bt.NewBinaryTreeNode(1)
	btVar.SetLeft(bt.NewBinaryTreeNode(2))
	btVar.SetRight(bt.NewBinaryTreeNode(3))

	fmt.Println("Root Node: ", btVar.GetData())
	fmt.Println("Left Node: ", btVar.GetLeft().GetData())
	fmt.Println("Right Node: ", btVar.GetRight().GetData())

	fmt.Println("------------------------")

	fmt.Println("Is Root Node a Leaf Node: ", btVar.IsLeaf())
	fmt.Println("Is Root Node a Full Node: ", btVar.IsFull())

	fmt.Println("Is Left Node a Leaf Node: ", btVar.GetLeft().IsLeaf())
	fmt.Println("Is Left Node a Leaf Node: ", btVar.GetLeft().IsLeaf())

	fmt.Println("------------------------")

	btVar.GetLeft().SetLeft(bt.NewBinaryTreeNode(4))

	fmt.Println("Left -> Left Node: ", btVar.GetLeft().GetLeft().GetData())
	fmt.Println("Is Left Node a Leaf Node: ", btVar.GetLeft().IsLeaf())
	fmt.Println("Is Left Node a Full Node: ", btVar.GetLeft().IsFull())

	fmt.Println("------------------------")

	btVar.GetLeft().SetRight(bt.NewBinaryTreeNode(5))

	fmt.Println("Left -> Right Node: ", btVar.GetLeft().GetRight().GetData())
	fmt.Println("Is Left Node a Leaf Node: ", btVar.GetLeft().IsLeaf())
	fmt.Println("Is Left Node a Full Node: ", btVar.GetLeft().IsFull())

	fmt.Println("------------------------")

	btVar.GetRight().SetLeft(bt.NewBinaryTreeNode(6))
	btVar.GetRight().SetRight(bt.NewBinaryTreeNode(7))

	fmt.Println("Right -> Left Node: ", btVar.GetRight().GetLeft().GetData())
	fmt.Println("Right -> Right Node: ", btVar.GetRight().GetRight().GetData())

	fmt.Println("Is Right Node a Leaf Node: ", btVar.GetRight().IsLeaf())
	fmt.Println("Is Right Node a Full Node: ", btVar.GetRight().IsFull())

	fmt.Println("------------------------")

	fmt.Println("InOrder Traversal: ")
	btVar.InOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("PreOrder Traversal: ")
	btVar.PreOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("PostOrder Traversal: ")
	btVar.PostOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("Found 4: ", btVar.Search(4))
	fmt.Println("Found 5: ", btVar.Search(5))
	fmt.Println("Found 9: ", btVar.Search(9))

	fmt.Println("------------------------")

	fmt.Println("Found 2: ", btVar.Search(2))
	fmt.Println("Deleting 2: ", btVar.Delete(2, nil))
	fmt.Println("Found 2: ", btVar.Search(2))

	fmt.Println("------------------------")

	fmt.Println("InOrder Traversal: ")
	btVar.InOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("Found 1: ", btVar.Search(1))
	fmt.Println("Deleting 1: ", btVar.Delete(1, nil))
	fmt.Println("Found 1: ", btVar.Search(1))

	fmt.Println("------------------------")

	fmt.Println("Root Node: ", btVar.GetData())
	fmt.Println("Left Node: ", btVar.GetLeft().GetData())
	fmt.Println("Right Node: ", btVar.GetRight().GetData())

	fmt.Println("------------------------")

	fmt.Println("Left -> Left Node: ", btVar.GetLeft().GetLeft().GetData())
	// fmt.Println("Left -> Right Node: ", btVar.GetLeft().GetRight().GetData())

	// fmt.Println("Right -> Left Node: ", btVar.GetRight().GetLeft().GetData())
	fmt.Println("Right -> Right Node: ", btVar.GetRight().GetRight().GetData())

	fmt.Println("------------------------")

	fmt.Println("InOrder Traversal: ")
	btVar.InOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("Implmeneing Binary Search Tree...")

	bstVar := bst.NewBinarySearchTreeNode(10)
	bstVar.Insert(5)
	bstVar.Insert(15)
	bstVar.Insert(2)

	fmt.Println("Root Node: ", bstVar.GetData())
	fmt.Println("Left Node: ", bstVar.GetLeft().GetData())
	fmt.Println("Right Node: ", bstVar.GetRight().GetData())
	fmt.Println("Left -> Left Node: ", bstVar.GetLeft().GetLeft().GetData())

	fmt.Println("------------------------")

	fmt.Println("InOrder Traversal: ")
	bstVar.InOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("PreOrder Traversal: ")
	bstVar.PreOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("PostOrder Traversal: ")
	bstVar.PostOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("Found 5: ", bstVar.Search(5))
	fmt.Println("Found 15: ", bstVar.Search(15))
	fmt.Println("Found 9: ", bstVar.Search(9))

	fmt.Println("------------------------")

	fmt.Println("Implmenting Ternary Tree...")

	fmt.Println("------------------------")

	ttVar := tt.NewTernaryTreeNode(30)
	ttVar.SetLeft(tt.NewTernaryTreeNode(5))
	ttVar.SetMid(tt.NewTernaryTreeNode(11))
	ttVar.SetRight(tt.NewTernaryTreeNode(63))

	fmt.Println("Root Node: ", ttVar.GetData())
	fmt.Println("Left Node: ", ttVar.GetLeft().GetData())
	fmt.Println("Mid Node: ", ttVar.GetMid().GetData())
	fmt.Println("Right Node: ", ttVar.GetRight().GetData())

	fmt.Println("------------------------")

	ttVar.GetLeft().SetLeft(tt.NewTernaryTreeNode(1))
	ttVar.GetLeft().SetMid(tt.NewTernaryTreeNode(4))
	ttVar.GetLeft().SetRight(tt.NewTernaryTreeNode(8))

	fmt.Println("Left -> Left Node: ", ttVar.GetLeft().GetLeft().GetData())
	fmt.Println("Left -> Mid Node: ", ttVar.GetLeft().GetMid().GetData())
	fmt.Println("Left -> Right Node: ", ttVar.GetLeft().GetRight().GetData())

	fmt.Println("------------------------")

	ttVar.GetMid().SetLeft(tt.NewTernaryTreeNode(6))
	ttVar.GetMid().SetMid(tt.NewTernaryTreeNode(7))
	ttVar.GetMid().SetRight(tt.NewTernaryTreeNode(15))

	fmt.Println("Mid -> Left Node: ", ttVar.GetMid().GetLeft().GetData())
	fmt.Println("Mid -> Mid Node: ", ttVar.GetMid().GetMid().GetData())
	fmt.Println("Mid -> Right Node: ", ttVar.GetMid().GetRight().GetData())

	fmt.Println("------------------------")

	ttVar.GetRight().SetLeft(tt.NewTernaryTreeNode(31))
	ttVar.GetRight().SetMid(tt.NewTernaryTreeNode(55))
	ttVar.GetRight().SetRight(tt.NewTernaryTreeNode(65))

	fmt.Println("Right -> Left Node: ", ttVar.GetRight().GetLeft().GetData())
	fmt.Println("Right -> Mid Node: ", ttVar.GetRight().GetMid().GetData())
	fmt.Println("Right -> Right Node: ", ttVar.GetRight().GetRight().GetData())

	fmt.Println("------------------------")

	fmt.Println("InOrder Traversal: ")
	ttVar.InOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	// https://www.geeksforgeeks.org/create-doubly-linked-list-ternary-ree/
	fmt.Println("PreOrder Traversal (kind of doubly linked list): ")
	ttVar.PreOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("PostOrder Traversal: ")
	ttVar.PostOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("Found 7: ", ttVar.Search(7))
	fmt.Println("Found 15: ", ttVar.Search(15))
	fmt.Println("Found 9: ", ttVar.Search(9))

	fmt.Println("------------------------")

	fmt.Println("Implmenting N-ary Tree...")

	nryVar := nry.NewNaryTreeNode(30)
	nryVar.AddChild(nry.NewNaryTreeNode(5))
	nryVar.AddChild(nry.NewNaryTreeNode(11))
	nryVar.AddChild(nry.NewNaryTreeNode(63))

	fmt.Println("Root Node: ", nryVar.GetData())
	fmt.Println("Child 1: ", nryVar.GetChildren()[0].GetData())
	fmt.Println("Child 2: ", nryVar.GetChildren()[1].GetData())
	fmt.Println("Child 3: ", nryVar.GetChildren()[2].GetData())

	fmt.Println("------------------------")

	nryVar.GetChildren()[0].AddChild(nry.NewNaryTreeNode(1))
	nryVar.GetChildren()[0].AddChild(nry.NewNaryTreeNode(4))
	nryVar.GetChildren()[0].AddChild(nry.NewNaryTreeNode(8))

	fmt.Println("Child 1 -> Child 1: ", nryVar.GetChildren()[0].GetChildren()[0].GetData())
	fmt.Println("Child 1 -> Child 2: ", nryVar.GetChildren()[0].GetChildren()[1].GetData())
	fmt.Println("Child 1 -> Child 3: ", nryVar.GetChildren()[0].GetChildren()[2].GetData())

	fmt.Println("------------------------")

	nryVar.GetChildren()[1].AddChild(nry.NewNaryTreeNode(6))
	nryVar.GetChildren()[1].AddChild(nry.NewNaryTreeNode(7))
	nryVar.GetChildren()[1].AddChild(nry.NewNaryTreeNode(15))

	fmt.Println("Child 2 -> Child 1: ", nryVar.GetChildren()[1].GetChildren()[0].GetData())
	fmt.Println("Child 2 -> Child 2: ", nryVar.GetChildren()[1].GetChildren()[1].GetData())
	fmt.Println("Child 2 -> Child 3: ", nryVar.GetChildren()[1].GetChildren()[2].GetData())

	fmt.Println("------------------------")

	nryVar.GetChildren()[2].AddChild(nry.NewNaryTreeNode(31))
	nryVar.GetChildren()[2].AddChild(nry.NewNaryTreeNode(55))
	nryVar.GetChildren()[2].AddChild(nry.NewNaryTreeNode(65))

	fmt.Println("Child 3 -> Child 1: ", nryVar.GetChildren()[2].GetChildren()[0].GetData())
	fmt.Println("Child 3 -> Child 2: ", nryVar.GetChildren()[2].GetChildren()[1].GetData())
	fmt.Println("Child 3 -> Child 3: ", nryVar.GetChildren()[2].GetChildren()[2].GetData())

	fmt.Println("------------------------")

	fmt.Println("Found 3: ", nryVar.Search(3))
	fmt.Println("Found 8: ", nryVar.Search(8))
	fmt.Println("Found 9: ", nryVar.Search(9))

	fmt.Println("------------------------")

	fmt.Println("InOrder Traversal: ")
	nryVar.InOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("PreOrder Traversal: ")
	nryVar.PreOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("PostOrder Traversal: ")
	nryVar.PostOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("Implmenting N-ary Tree with second approach...")

	nryVar2 := nrysib.NewTreeNode(30)
	nryVar2.AddChild(nrysib.NewTreeNode(5))
	nryVar2.AddChild(nrysib.NewTreeNode(11))
	nryVar2.AddChild(nrysib.NewTreeNode(63))

	fmt.Println("Root Node: ", nryVar2.GetData())
	fmt.Println("Child 1: ", nryVar2.GetChildren().GetData())
	fmt.Println("Child 2: ", nryVar2.GetChildren().GetSibling().GetData())
	fmt.Println("Child 3: ", nryVar2.GetChildren().GetSibling().GetSibling().GetData())

	fmt.Println("------------------------")

	nryVar2.GetChildren().AddChild(nrysib.NewTreeNode(1))
	nryVar2.GetChildren().AddChild(nrysib.NewTreeNode(4))
	nryVar2.GetChildren().AddChild(nrysib.NewTreeNode(8))

	fmt.Println("Child 1 -> Child 1: ", nryVar2.GetChildren().GetChildren().GetData())
	fmt.Println("Child 1 -> Child 2: ", nryVar2.GetChildren().GetChildren().GetSibling().GetData())
	fmt.Println("Child 1 -> Child 3: ", nryVar2.GetChildren().GetChildren().GetSibling().GetSibling().GetData())

	fmt.Println("------------------------")

	nryVar2.GetChildren().GetSibling().AddChild(nrysib.NewTreeNode(6))
	nryVar2.GetChildren().GetSibling().AddChild(nrysib.NewTreeNode(7))
	nryVar2.GetChildren().GetSibling().AddChild(nrysib.NewTreeNode(15))

	fmt.Println("Child 2 -> Child 1: ", nryVar2.GetChildren().GetSibling().GetChildren().GetData())
	fmt.Println("Child 2 -> Child 2: ", nryVar2.GetChildren().GetSibling().GetChildren().GetSibling().GetData())
	fmt.Println("Child 2 -> Child 3: ", nryVar2.GetChildren().GetSibling().GetChildren().GetSibling().GetSibling().GetData())

	fmt.Println("------------------------")

	nryVar2.GetChildren().GetSibling().GetSibling().AddChild(nrysib.NewTreeNode(31))
	nryVar2.GetChildren().GetSibling().GetSibling().AddChild(nrysib.NewTreeNode(55))
	nryVar2.GetChildren().GetSibling().GetSibling().AddChild(nrysib.NewTreeNode(65))

	fmt.Println("Child 3 -> Child 1: ", nryVar2.GetChildren().GetSibling().GetSibling().GetChildren().GetData())
	fmt.Println("Child 3 -> Child 2: ", nryVar2.GetChildren().GetSibling().GetSibling().GetChildren().GetSibling().GetData())
	fmt.Println("Child 3 -> Child 3: ", nryVar2.GetChildren().GetSibling().GetSibling().GetChildren().GetSibling().GetSibling().GetData())

	fmt.Println("------------------------")

	fmt.Println("Found 3: ", nryVar2.Search(3))
	fmt.Println("Found 8: ", nryVar2.Search(8))
	fmt.Println("Found 9: ", nryVar2.Search(9))

	fmt.Println("------------------------")

	fmt.Println("PreOrder Traversal: ")
	nryVar2.PreOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")

	fmt.Println("PostOrder Traversal: ")
	nryVar2.PostOrderTraversal()
	fmt.Println()

	fmt.Println("------------------------")
}
