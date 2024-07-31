package lc_0206_reverse_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList .
func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}

	var pre, cur, next *ListNode
	pre = head
	cur = head.Next

	head.Next = nil
	if cur != nil {
		next = cur.Next
	}

	for next != nil {
		cur.Next = pre

		pre = cur
		cur = next
		next = next.Next
	}

	cur.Next = pre

	return cur

}
