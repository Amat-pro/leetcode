package temp

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	var pre *ListNode
	var cur *ListNode

	pre = nil
	cur = head

	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}

	return pre

}
