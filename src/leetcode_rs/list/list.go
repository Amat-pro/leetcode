package list

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) Println() {
	cur := n
	for cur != nil {
		fmt.Print(" ", cur.Value)
		cur = cur.Next
	}
}

// removeElement 203 - 移除链表元素
func removeElement(head *Node, val int) *Node {
	virtualHead := &Node{Next: head}

	pre := virtualHead
	cur := head
	for cur != nil {
		if cur.Value != val {
			cur = cur.Next
			pre = pre.Next

			continue
		}

		next := cur.Next
		pre.Next = next
		cur.Next = nil

		cur = next
	}

	return virtualHead.Next

}

// reverseList 206 - 翻转链表
func reverseList(head *Node) *Node {
	pre := head
	cur := head.Next

	// 这里需把pre.next设置为nil, 否则 head -> n2 <- n3 <- n4 <- newHead
	//                                   <- n2 （导致循环链表）
	pre.Next = nil

	for cur != nil {

		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next

	}

	return pre
}

// swapPairs 24 - 两两交换链表中的节点
func swapPairs(head *Node) *Node {

	virtualHead := &Node{}
	virtualHead.Next = head

	pre := virtualHead
	cur := head

	for cur != nil && cur.Next != nil {
		// 交换cur 和 cur.Next
		next := cur.Next.Next

		pre.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = next

		pre = cur
		cur = next
	}

	return virtualHead.Next

}

// removeNthFromEnd 19 - 删除链表的倒数第N个节点
func removeNthFromEnd(head *Node, n int) *Node {
	// 双指针
	// fast先走n个，当fast.Next == nil 时，slow.Next是需要删除的节点

	virtualHead := &Node{}
	virtualHead.Next = head

	slow := virtualHead
	fast := virtualHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	next := slow.Next.Next
	slow.Next.Next = nil
	slow.Next = next

	return virtualHead.Next

}

// getIntersectionNode 02.07 - 链表相交
func getIntersectionNode(headA, headB *Node) *Node {
	lenA := 0
	lenB := 0

	tempNode := headA
	for tempNode != nil {
		lenA++
		tempNode = tempNode.Next
	}
	tempNode = headB
	for tempNode != nil {
		lenB++
		tempNode = tempNode.Next
	}

	// headA为长链表
	if lenA < lenB {
		lenA, lenB = lenB, lenA
		headA, headB = headB, headA
	}

	n := lenA - lenB
	for i := 0; i < n; i++ {
		headA = headA.Next
	}

	fmt.Println("--> ", n, headA.Value, headB.Value)

	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil

}

// detectCycle 142 - 环形链表2
func detectCycle(head *Node) *Node {

	// 	快慢指针
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			temp := head
			for temp != slow {
				temp = temp.Next
				slow = slow.Next
			}
			return slow
		}

	}

	return nil

}
