package list

import (
	"fmt"
	"testing"
)

func Test_removeElements(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	head := removeElement(n1, 1)
	head.Println()
}

func Test_reverseList(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	head := reverseList(n1)
	head.Println()
}

func Test_swapPairs(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n6 := &Node{Value: 6}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	head := swapPairs(n1)
	head.Println()
}

func Test_removeNthFromEnd(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n6 := &Node{Value: 6}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	head := removeNthFromEnd(n1, 6) // 移除n1
	head.Println()
}

func Test_getIntersectionNode(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n6 := &Node{Value: 6}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6

	b1 := &Node{Value: 11}
	b1.Next = n4

	node := getIntersectionNode(n1, b1)
	if node != nil {
		fmt.Println("getIntersectionNode result is: 4 ", node.Value) // OUTPUT this
	} else {
		fmt.Println("no node")
	}
}

func Test_detectCycle(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n6 := &Node{Value: 6}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n2

	node := detectCycle(n1)
	if node != nil {
		fmt.Println("detectCycle result is: 2 ", node.Value)
	}
}
