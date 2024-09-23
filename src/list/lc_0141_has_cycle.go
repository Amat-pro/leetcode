package list

// hasCycle 环形链表 - 快慢指针
func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	var slow, fast *ListNode
	slow = head
	fast = head
	for fast != nil && fast.Next != nil {

		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}

	}

	return false

}