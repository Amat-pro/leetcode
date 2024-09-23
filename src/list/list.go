package list

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) print() {
	var cur *ListNode
	cur = l 
	for cur != nil {
		fmt.Printf("%d ",  cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}