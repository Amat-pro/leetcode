package main

import (
	"container/heap"
	"fmt"
	"testing"
)

// 定义一个最大堆的类型
type MaxHeap []int

// 实现 heap.Interface 的方法

// Len 返回堆中元素的个数
func (h MaxHeap) Len() int {
	return len(h)
}

// Less 定义元素的比较规则，这里实现最大堆
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap 交换堆中的两个元素
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 向堆中添加元素
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 从堆中移除并返回最大的元素
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestHeap(t *testing.T) {
	// 创建一个最大堆
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	// 向堆中添加元素
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 2)
	heap.Push(maxHeap, 4)
	heap.Push(maxHeap, 1)

	// 输出堆中的元素
	fmt.Println("MaxHeap:", *maxHeap)

	// 从堆中弹出最大的元素
	fmt.Println("Pop:", heap.Pop(maxHeap))

	// 再次输出堆中的元素
	fmt.Println("MaxHeap after Pop:", *maxHeap)
}
