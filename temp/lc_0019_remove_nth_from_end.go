package temp

func removeNthFromEnd(head *Node, n int) *Node {

	vitualHead := &Node{Val: -1} // 使用虚拟节点，删除的节点有可能是跟节点！！
	vitualHead.Next = head

	slow, fast := vitualHead, vitualHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return vitualHead.Next


}