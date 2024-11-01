package temp

func detectCycle(head *ListNode) *ListNode {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {

			for head != slow {
				head = head.Next
				slow = slow.Next
			}

			return slow

		}
	}

	return nil

}
