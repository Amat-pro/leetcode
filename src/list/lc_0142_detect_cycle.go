package list 

// detectCycle 环形链表II
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {

		fast = fast.Next.Next
		slow = slow.Next

		// 快慢指针，环形相遇后从head出发，slow和head每次移动一个节点，
		// slow和head的相遇点就是环形入口
		if slow == fast {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}

			return slow
		}

	}

	return nil

}