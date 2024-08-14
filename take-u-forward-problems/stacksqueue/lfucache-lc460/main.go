package main

import (
	"fmt"
)

// https://leetcode.com/problems/lfu-cache/description/
type LFUCache struct {
	nodeMap  map[int]*Node
	listMap  map[int]*List
	capacity int
	min      int
}

type List struct {
	head *Node
	tail *Node
	size int
}

func (l *List) PushNode(node *Node) {
	node.prev = l.tail.prev
	node.next = l.tail
	l.tail.prev.next = node
	l.tail.prev = node
	l.size++
}

func (l *List) DeleteNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
}

func makeList() *List {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &List{
		head: head,
		tail: tail,
	}
}

type Node struct {
	prev, next     *Node
	key, val, freq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		nodeMap:  make(map[int]*Node),
		listMap:  make(map[int]*List),
		capacity: capacity,
		min:      0,
	}
}

func (lfu *LFUCache) Get(key int) int {
	node, ok := lfu.nodeMap[key]
	if !ok {
		return -1
	}
	list, ok := lfu.listMap[node.freq]
	if ok {
		list.DeleteNode(node)
	}

	node.freq++
	nextList, nextOk := lfu.listMap[node.freq]
	if !nextOk {
		nextList = makeList()
	}
	nextList.PushNode(node)
	lfu.listMap[node.freq] = nextList
	if list.size == 0 && lfu.min == node.freq-1 {
		lfu.min++
	}
	return node.val
}

func (lfu *LFUCache) Put(key int, value int) *LFUCache {
	if lfu.capacity == 0 {
		return nil
	}

	node, ok := lfu.nodeMap[key]
	if ok {
		node.val = value
		lfu.Get(key)
		return nil
	}

	// exceed maximum capacity
	if len(lfu.nodeMap) == lfu.capacity {
		minList := lfu.listMap[lfu.min]
		leastFrequencyNode := minList.head.next
		minList.DeleteNode(leastFrequencyNode)
		delete(lfu.nodeMap, leastFrequencyNode.key)
	}

	node = &Node{
		key:  key,
		val:  value,
		freq: 1,
	}

	lfu.min = 1
	list, ok := lfu.listMap[node.freq]
	if !ok {
		list = makeList()
	}
	list.PushNode(node)
	lfu.listMap[node.freq] = list
	lfu.nodeMap[key] = node
	return nil
}

func main() {
	obj := Constructor(2)

	fmt.Println(
		"null", obj.Put(1, 1), obj.Put(2, 2),
		obj.Get(1), obj.Put(3, 3), obj.Get(2), obj.Get(3),
		obj.Put(4, 4), obj.Get(1), obj.Get(3), obj.Get(4),
	)

	// obj := Constructor(3)
	// fmt.Println(
	// 	"null", obj.Put(2, 2), obj.Put(1, 1),
	// 	obj.Get(2), obj.Get(1), obj.Get(2),
	// 	obj.Put(3, 3), obj.Put(4, 4),
	// 	obj.Get(3), obj.Get(2), obj.Get(1), obj.Get(4),
	// )
}
