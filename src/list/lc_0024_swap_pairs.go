package list

// swapPairs 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	dummyHead := &ListNode{}
	dummyHead.Next = head

	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {

		temp1 := cur.Next

		cur.Next = cur.Next.Next
		temp1.Next = cur.Next.Next
		cur.Next.Next = temp1

		cur = cur.Next.Next

	}

	return dummyHead.Next

}