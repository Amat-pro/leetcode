package cache

import "fmt"

// LRUCache 实现

type Node struct {
	Key   int
	Value int
	Pre   *Node
	Next  *Node
}

type LRUCache struct {
	capacity  int
	length    int
	head      *Node
	tail      *Node
	keyToNode map[int]*Node
}

func New(capacity int) *LRUCache {
	if capacity < 0 {
		panic("invalid capacity")
	}

	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Pre = head

	return &LRUCache{
		capacity:  capacity,
		length:    0,
		head:      head,
		tail:      tail,
		keyToNode: make(map[int]*Node, capacity),
	}
}

func (cache *LRUCache) Get(key int) int {
	node, ok := cache.keyToNode[key]
	if ok {
		cache.moveNext(cache.head, node)
		return node.Value
	}

	return -1
}

func (cache *LRUCache) Put(key, value int) {
	node, ok := cache.keyToNode[key]
	if ok {
		node.Value = value
		cache.moveNext(cache.head, node)
		return
	}

	if cache.capacity == cache.length {
		delete(cache.keyToNode, cache.tail.Pre.Key) // 删除tail.Pre节点
		cache.removeNode(cache.tail.Pre)
		cache.length--
	}

	addNode := &Node{
		Key:   key,
		Value: value,
	}
	cache.keyToNode[key] = addNode
	cache.moveNext(cache.head, addNode)
	cache.length++

}

func (cache *LRUCache) moveNext(node *Node, addNode *Node) {
	if addNode.Next != nil && addNode.Pre != nil { // 移除addNode
		pre := addNode.Pre
		next := addNode.Next
		addNode.Pre = nil
		addNode.Next = nil

		pre.Next = next
		next.Pre = pre
	}

	// 添加到node的next位置
	next := node.Next
	node.Next = addNode
	addNode.Pre = node
	addNode.Next = next
	next.Pre = addNode

}

func (cache *LRUCache) removeNode(node *Node) {
	pre := node.Pre
	next := node.Next
	node.Pre = nil
	node.Next = nil

	pre.Next = next
	next.Pre = pre
}

func (cache *LRUCache) Print() {
	cur := cache.head.Next
	for cur != cache.tail {
		fmt.Print("  ", cur.Key)
		cur = cur.Next
	}
	fmt.Println()
}
