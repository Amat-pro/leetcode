package temp

func reverseList(head *Node) *Node {

	var pre *Node
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}

	return pre

}
