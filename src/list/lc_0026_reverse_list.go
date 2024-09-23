package list

// reverseList 反转链表
func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	var cur *ListNode
	var pre *ListNode

	cur = head.Next
	pre = head
	pre.Next = nil

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur

		cur = next
	}

	return pre

} 