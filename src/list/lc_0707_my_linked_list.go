package list

// MyLinkedList 设计链表
type MyLinkedList struct {

	Head *ListNode
	Len int

}


func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: &ListNode{},
	}
}

// Get 获取第index个节点
func (l *MyLinkedList) Get(index int) int {

	if index < 0 || index > l.Len -1 {
		return -1
	}

	cur := l.Head

	for index > 0 {
		cur = cur.Next
		index --
	}

	return cur.Next.Val
}

// AddAtHead 头部插入节点
func (l *MyLinkedList) AddAtHead(val int)  {
	newNode := &ListNode{
		Val: val,
	}
	newNode.Next = l.Head.Next
	l.Head.Next = newNode

	l.Len++

}

// AddAtTail 尾部插入节点
func (l *MyLinkedList) AddAtTail(val int)  {

	tail := l.Head
	for tail.Next != nil {
		tail = tail.Next
	}

	newNode := ListNode{Val: val}
	tail.Next = &newNode

	l.Len++

}

// AddAtIndex 第n个节点前插入节点
func (l *MyLinkedList) AddAtIndex(index int, val int)  {

	if index < 0 || index > l.Len -1 {
		panic("invalid index")
	}

	newNode := &ListNode{Val: val}
	cur := l.Head
	for index > 0 {
		cur = cur.Next
		index --
	}
	newNode.Next = cur.Next
	cur.Next = newNode
	l.Len++

}

// DeleteAtIndex 删除第index个节点
func (l *MyLinkedList) DeleteAtIndex(index int)  {

	if index < 0 || index > l.Len -1 {
		panic("invalid index")
	}

	cur := l.Head
	for index > 0 {
		cur = cur.Next
		index --
	}

	if cur.Next != nil {
		cur.Next = cur.Next.Next
	} else {
		cur.Next = nil
	}

	l.Len--

}