package lru

import "fmt"


type LRUCache2 struct {
	capacity int
	length int
	head *Node
	keyToValue map[int]*Node
}

type Node struct {
	Pre *Node
	Next *Node
	Key int
	Value int
}

func Constructor2(capacity int) LRUCache2 {
	if capacity <= 0 {
		panic("invalid capacity")
	}

	head := Node{}
	tail := Node{}
	head.Pre = &tail

	return LRUCache2{
		capacity: capacity,
		length: 0,
		head: &head,
		keyToValue: make(map[int]*Node, capacity),
	}
}

func (cache *LRUCache2) Get(key int) int {

	v, ok := cache.keyToValue[key]
	if !ok {
		return -1
	}

	cache.moveNNext(cache.head, v)
	return v.Value

}

func (cache *LRUCache2) Put(key int, value int) {

	v, ok := cache.keyToValue[key]
	if ok {
		cache.moveNNext(cache.head, v)
		return
	}

	// 需要新增
	// 如果满了，移除一个
	if cache.length == cache.capacity {
		cache.removeTailPre()
	}

	cache.addHadNext(key, value)
}

// addHeadNext 添加kv到head的next
func (cache *LRUCache2) addHadNext(key int, value int) {
	cache.length ++
	n := &Node{
		Key: key,
		Value: value,
	}
	cache.keyToValue[key] = n

	temp := cache.head.Next
	cache.head.Next = n
	n.Next = temp
	n.Pre = cache.head
	temp.Pre = n

}

// removeTailPre 移除tail的pre节点
func (cache *LRUCache2) removeTailPre() {

	if cache.length == 0 {
		return
	}

	tail := cache.head.Pre
	remove := tail.Pre
	removePre := remove.Pre
	remove.Pre = nil
	remove.Next = nil
	removePre.Next = tail
	tail.Pre = removePre

}

// moveNNext 移动moved到n的next
func (cache *LRUCache2) moveNNext(n *Node, moved *Node) {
	moved.Pre.Next = moved.Next
	moved.Next.Pre = moved.Pre

	n.Next.Pre = moved
	moved.Next = n.Next

	n.Next = moved
	moved.Pre = n

}

func (cache *LRUCache2) Print() {

	node := cache.head.Next

	for node != nil {
		fmt.Printf("node, key: %d, value: %d\n", node.Key, node.Value)
		node = node.Next
	}

}