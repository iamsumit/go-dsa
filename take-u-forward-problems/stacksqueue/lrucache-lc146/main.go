package main

import "fmt"

// https://takeuforward.org/data-structure/implement-lru-cache/
type Node struct {
	prev, next *Node
	key, value int
}

type LRUCache struct {
	head, tail *Node
	cache      map[int]*Node
	capacity   int
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		head:     head,
		tail:     tail,
		cache:    make(map[int]*Node),
		capacity: capacity,
	}
}

func (lc *LRUCache) Get(key int) int {
	node, exists := lc.cache[key]
	if !exists {
		return -1
	}

	lc.remove(node)
	lc.insert(node)
	return node.value
}

func (lc *LRUCache) Put(key int, value int) {
	if node, exists := lc.cache[key]; exists {
		lc.remove(node)
	} else if len(lc.cache) == lc.capacity {
		lc.remove(lc.tail.prev)
	}

	lc.insert(&Node{key: key, value: value})
}

func (lc *LRUCache) remove(node *Node) {
	delete(lc.cache, node.key)
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lc *LRUCache) insert(node *Node) {
	lc.cache[node.key] = node
	node.next = lc.head.next
	node.next.prev = node
	lc.head.next = node
	node.prev = lc.head
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println("obj.Put(1, 1)")
	fmt.Println("obj.Put(2, 2)")
	fmt.Println("obj.Get(1)", obj.Get(1))
	obj.Put(3, 3)
	fmt.Println("obj.Put(3, 3)")
	fmt.Println("obj.Get(1)", obj.Get(1))
	fmt.Println("obj.Get(2)", obj.Get(2))
	obj.Put(4, 4)
	fmt.Println("obj.Put(4, 4)")
	fmt.Println("obj.Get(1)", obj.Get(1))
	fmt.Println("obj.Get(3)", obj.Get(3))
	fmt.Println("obj.Get(4)", obj.Get(4))
}
