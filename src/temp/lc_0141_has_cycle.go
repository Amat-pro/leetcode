package temp

func hasCycle(head *Node) bool {
	// 快慢指针
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}

	}

	return false

}

func hasCycleII(head *Node) *Node {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next 

		if slow == fast {
			ptr := head
			for {
				slow = slow.Next
				ptr = ptr.Next

				if slow == ptr {
					return slow
				}
			}
		}
	}

	return nil
}