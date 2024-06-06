package ll

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Insert(data int) {
	node := &Node{Data: data}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	l.Tail.Next = node
	node.Prev = l.Tail
	l.Tail = node
}

func (l *LinkedList) Delete(data int) {
	if l.Head == nil {
		return
	}

	if l.Head.Data == data {
		l.Head = l.Head.Next
		l.Head.Prev = nil
		return
	}

	if l.Tail.Data == data {
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
		return
	}

	for n := l.Head; n != nil; n = n.Next {
		if n.Data == data {
			n.Prev.Next = n.Next
			n.Next.Prev = n.Prev
			return
		}
	}
}

func (l *LinkedList) Search(data int) *Node {
	for n := l.Head; n != nil; n = n.Next {
		if n.Data == data {
			return n
		}
	}

	return nil
}

func (l *LinkedList) Reverse() {
	if l.Head == nil {
		return
	}

	var prev *Node
	curr := l.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		curr.Prev = next
		prev = curr
		curr = next
	}

	l.Head = prev
}
