package temp

import (
	"fmt"
	"testing"
)

func Test_findLengthOfLCIS(t *testing.T) {
	fmt.Println("最长连续递增序列： ", findLengthOfLCIS([]int{1, 3, 5, 4, 7})) // 3
}

func Test_longestConsecutive(t *testing.T) {
	fmt.Println("最长递增序列：", longestConsecutive([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
}

func Test__getInterectionNode(t *testing.T) {
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	n3 := ListNode{Val: 3}
	n4 := ListNode{Val: 4}
	n5 := ListNode{Val: 5}

	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5

	a1 := ListNode{Val: -1}
	a1.Next = &n4

	fmt.Println("链表相交，result is: ", getIntersectionNode(&n1, &a1).Val)

}

func Test_reverseList(t *testing.T) {
	n1 := ListNode{Val: 1}
	n2 := ListNode{Val: 2}
	n3 := ListNode{Val: 3}
	n4 := ListNode{Val: 4}
	n5 := ListNode{Val: 5}

	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5

	headNew := reverseList(&n1)
	for headNew != nil {
		fmt.Println("val: ", headNew.Val)
		headNew = headNew.Next

	}
}
