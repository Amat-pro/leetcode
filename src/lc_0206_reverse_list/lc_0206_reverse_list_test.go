package lc_0206_reverse_list

import (
	"fmt"
	"testing"
)

func TestLc0206ReverseList(t *testing.T) {
	n1 := ListNode{
		Val: 1,
	}
	n2 := ListNode{
		Val: 2,
	}
	n3 := ListNode{
		Val: 3,
	}
	n4 := ListNode{
		Val: 4,
	}
	n5 := ListNode{
		Val: 5,
	}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5

	reverseList := reverseList(&n1)
	node := reverseList
	for node != nil {
		fmt.Println("node value: ", node.Val)
		node = node.Next
	}

}

func TestLc0206ReverseList_02(t *testing.T) {
	n1 := ListNode{
		Val: 1,
	}
	n2 := ListNode{
		Val: 2,
	}

	n1.Next = &n2

	reverseList := reverseList(&n1)
	node := reverseList
	for node != nil {
		fmt.Println("node value: ", node.Val)
		node = node.Next
	}

}

func TestLc0206ReverseList_03(t *testing.T) {
	n1 := ListNode{
		Val: 1,
	}

	reverseList := reverseList(&n1)
	node := reverseList
	for node != nil {
		fmt.Println("node value: ", node.Val)
		node = node.Next
	}

}

func TestLc0206ReverseList_v2(t *testing.T) {
	n1 := ListNode{
		Val: 1,
	}
	n2 := ListNode{
		Val: 2,
	}
	n3 := ListNode{
		Val: 3,
	}
	n4 := ListNode{
		Val: 4,
	}
	n5 := ListNode{
		Val: 5,
	}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5

	reverseList := reverseList_v2(&n1)
	node := reverseList
	for node != nil {
		fmt.Println("node value: ", node.Val)
		node = node.Next
	}

}
