package list

import (
	"fmt"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	
	n1.print()
	mewList := removeElements(n1,3)
	mewList.print()

}

func TestMyLinkedList(t *testing.T) {
	l := Constructor()

	l.AddAtHead(1)  
	l.Head.print()
	l.AddAtTail(3)
	l.Head.print()
	l.AddAtIndex(1,2)
	l.Head.print()
	l.Get(1)
	l.Head.print()
	l.DeleteAtIndex(1)
	l.Head.print()
	l.Get(1)

}

func TestReverseList(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	
	n1.print()
	mewList := reverseList(n1)
	mewList.print()
}

func TestSwapPairs(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	n1.print()
	newHead := swapPairs(n1)
	newHead.print() // 2 1 4 3

}

func TestRemoveNthFromEnd(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	n1.print()
	newNode := removeNthFromEnd(n1, 5)
	newNode.print()
}

func TestGetIntersectionNode(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	a1 := &ListNode{Val: 1}
	a2 := &ListNode{Val: 2}
	a1.Next = a2
	a2.Next = n3

	fmt.Println("链表相交节点值: ", getIntersectionNode(a1, n1).Val) // 3
}

func TestHasCycle(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	fmt.Println("环形链表 result is: ", hasCycle(n1)) // false
	n5.Next = n2
	fmt.Println("环形链表 result is: ", hasCycle(n1)) // true
}

func TestDetectCycle(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n3
	fmt.Println("环形链表II result is: ", detectCycle(n1).Val) // 3
}