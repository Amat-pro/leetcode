package tree

import (
	"fmt"
	"testing"
)

func TestTraversal(t *testing.T) {

	//       1
	// 	 2         3
	// 4    5

	n1 := &TreeNode {
		Val: 1,
	}
	n2 := &TreeNode {
		Val: 2,
	}
	n3 := &TreeNode {
		Val: 3,
	}
	n4 := &TreeNode {
		Val: 4,
	}
	n5 := &TreeNode {
		Val: 5,
	}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5

	// 1 2 4 5 3
	fmt.Println("pre order result: ", preOrderTraversal(n1))
	fmt.Println("pre order result: ", preOrderTraversal_V2(n1))
	// 4 2 5 1 3
	fmt.Println("in order result: ", inOrderTraversal(n1))
	// 4 5 2 3 1
	fmt.Println("post order result: ", postOrderTraversal(n1))


}