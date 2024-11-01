package temp

import (
	"fmt"
	"testing"
)

func Test_lc_0739_daily_temperatures(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := dailyTemperatures(temperatures)
	fmt.Println("0739_daily_temperatures(每日温度) result is: ", result)
}

func Test_lc_0042_trap(t *testing.T) {
	h := []int{4, 2, 0, 3, 2, 5}
	result := trap(h)
	fmt.Println("0042_trap(接雨水) result is: ", result)
}

func Test_lc_0160_getInterectionNode(t *testing.T) {
	n1 := Node{Val: 1}
	n2 := Node{Val: 2}
	n3 := Node{Val: 3}
	n4 := Node{Val: 4}
	n5 := Node{Val: 5}

	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5

	a1 := Node{Val: -1}
	a1.Next = &n4

	fmt.Println("链表相交，result is: ", getInterectionNode(&n1, &a1).Val)

}

func Test_lc_0019_removeNthFromEnd(t *testing.T) {
	n1 := Node{Val: 1}
	n2 := Node{Val: 2}
	n3 := Node{Val: 3}
	n4 := Node{Val: 4}
	n5 := Node{Val: 5}

	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	n4.Next = &n5

	headNew := removeNthFromEnd(&n1, 2)
	for headNew != nil {
		fmt.Println("val: ", headNew.Val)
		headNew = headNew.Next
	}
}

func Test_lc_0015_threeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println("threeSum(三数之和), result is: ", result)
}

func Test_lc_0206_reverseList(t *testing.T) {
	n1 := Node{Val: 1}
	n2 := Node{Val: 2}
	n3 := Node{Val: 3}
	n4 := Node{Val: 4}
	n5 := Node{Val: 5}

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

func Test_lc_0244_reverseString(t *testing.T) {

	s := "hello"
	sB := []byte(s)
	reverseString(sB)
	fmt.Println("反转字符串,result is: ", string(sB))

}

func Test_lc_0509_fib(t *testing.T) {
	fmt.Println("fib, result is: ", fib(3))
}

func Test_lc_0459_repeatedSubStringPattern(t *testing.T) {
	fmt.Println("repeatedSubStringPattern, result is: ", repeatedSubStringPattern("abcabcabcabc"))
}

func Test_lc_0131_partition(t *testing.T) {
	fmt.Println("分割回文串：result is: ", partition("aab"))
}
