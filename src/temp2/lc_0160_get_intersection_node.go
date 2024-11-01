package temp

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	lenA := 0
	lenB := 0

	var tempNode *ListNode
	tempNode = headA
	for tempNode != nil {
		lenA++
		tempNode = tempNode.Next
	}
	tempNode = headB
	for tempNode != nil {
		lenB++
		tempNode = tempNode.Next
	}

	if lenA < lenB {
		headA, headB = headB, headA
		lenA, lenB = lenB, lenA
	}

	var nodeA *ListNode
	var nodeB *ListNode

	nodeA = headA
	nodeB = headB
	for i := 0; i < lenA-lenB; i++ {
		nodeA = nodeA.Next
	}

	for nodeA != nil {
		if nodeA == nodeB {
			return nodeA
		}
		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}

	return nil

}
