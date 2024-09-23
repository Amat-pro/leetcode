package list

// getIntersectionNode 链表相交
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    aLen := 0
	bLen := 0

	var temp *ListNode
	temp = headA
	for temp != nil {
		aLen ++
		temp = temp.Next
	}
	temp = headB
	for temp != nil {
		bLen ++
		temp = temp.Next
	}

	// headA 长链表
	// headB 短链表
	if aLen < bLen {
		headA, headB = headB, headA
	}

	var curA, curB *ListNode
	curA = headA
	curB = headB
	for i := 0; i < aLen - bLen; i++ {
		curA = curA.Next
	}

	for curA != curB {
		curA = curA.Next
		curB = curB.Next
	}

	return curA

}