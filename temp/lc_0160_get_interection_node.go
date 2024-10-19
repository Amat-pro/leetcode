package temp

func getInterectionNode(headA, headB *Node) *Node {

	lenA := 0
	lenB := 0
	tempNode := headA
	for tempNode != nil {
		lenA++
		tempNode = tempNode.Next
	}
	tempNode2 := headB
	for tempNode2 != nil {
		lenB++
		tempNode2 = tempNode2.Next
	}

	length := lenA - lenB
	// 让A代表长链表
	if lenA < lenB {
		headA, headB = headB, headA
		length = lenB - lenA
	}

	curA := headA	
	curB := headB
	for i := 0; i < length; i++ {
		curA = curA.Next
	}

	for curA != nil {
		curA = curA.Next
		curB = curB.Next

		if curA == curB {
			return curA
		}
	}

	return nil

}
