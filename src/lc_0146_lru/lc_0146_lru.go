package lru

import "fmt"

type LRUCache struct {
	head *Entry // 虚拟头节点 使用指针，避免构造时发生复制导致地址变化
	tail *Entry // 虚拟尾节点

	entries map[int]*Entry // map结构

	capacity int
	len      int
}

func Constructor(capacity int) LRUCache {

	head := Entry{}
	tail := Entry{
		Key: -1,
	}

	head.Next = &tail
	tail.Pre = &head

	return LRUCache{
		head: &head, // 使用指针，避免构造时发生复制导致地址变化
		tail: &tail,

		entries:  make(map[int]*Entry, capacity),
		capacity: capacity,
		len:      0,
	}

}

func (cache *LRUCache) Get(key int) int {
	val, ok := cache.entries[key]
	if !ok {
		return -1
	}
	cache.moveToAt(val, cache.head)
	return val.Value
}

func (cache *LRUCache) Put(key, value int) {
	val, ok := cache.entries[key]
	if ok {
		val.Value = value
		cache.moveToAt(val, cache.head)
		return
	}

	if cache.len > 0 && cache.len == cache.capacity {
		cache.remove(cache.tail.Pre)
	}

	entry := Entry{
		Key:   key,
		Value: value,
	}

	cache.insertToFront(&entry)
}

func (cache *LRUCache) Print() {

	node := cache.head.Next

	for node != nil {

		if node == cache.tail {
			return
		}

		fmt.Printf("node, key: %d, value: %d\n", node.Key, node.Value)
		node = node.Next
	}

}

// move move entry to at's next position
func (cache *LRUCache) moveToAt(entry, at *Entry) {

	if entry == at {
		return
	}

	entry.Pre.Next = entry.Next
	entry.Next.Pre = entry.Pre

	entry.Pre = at
	entry.Next = at.Next

	// at.Next = entry
	entry.Pre.Next = entry
	entry.Next.Pre = entry

}

// remove rmove entry from cache
func (cache *LRUCache) remove(entry *Entry) {
	pre := entry.Pre
	next := entry.Next

	pre.Next = next
	next.Pre = pre
	entry.Pre = nil
	entry.Next = nil

	delete(cache.entries, entry.Key)
	cache.len--

}

// insertFront insert entry at front
func (cache *LRUCache) insertToFront(entry *Entry) {
	head := cache.head
	next := head.Next

	head.Next = entry
	next.Pre = entry
	entry.Pre = head
	entry.Next = next

	cache.entries[entry.Key] = entry
	cache.len++

}

// Entry 双向链表
type Entry struct {
	Key, Value int

	Pre, Next *Entry
}
