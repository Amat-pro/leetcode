package list

// removeElements 移除链表元素
func removeElements(head *ListNode, val int) *ListNode {

	dummyHead := &ListNode {} // 虚拟头节点
	dummyHead.Next = head

	var cur *ListNode
	cur = dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}

		cur = cur.Next
	} 

	return dummyHead.Next
}