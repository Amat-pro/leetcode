 package list

 // removeNthFromEnd 移除链表的倒数第n个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummyHead := &ListNode{}
	dummyHead.Next = head

	var slow, fast *ListNode
	slow, fast = dummyHead, dummyHead

	num := 0  // [0, n+1)
	for num < n+1 && fast != nil {
		fast = fast.Next
		num++
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummyHead.Next

}