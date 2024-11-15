package tree

import (
	"fmt"
	"testing"
)

func Test_preOrderTraversal(t *testing.T) {
	root := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	root.Right = n2
	n2.Left = n3

	fmt.Println("preOrderTraversal result is: ", preOrderTraversal(root)) // [1,2,3]
	fmt.Println("preOrderTraversalWithStack result is: ", preOrderTraversalWithStack(root))
}

func Test_levelOrder(t *testing.T) {
	n3 := &TreeNode{Val: 3}
	n9 := &TreeNode{Val: 9}
	n20 := &TreeNode{Val: 20}
	n15 := &TreeNode{Val: 15}
	n7 := &TreeNode{Val: 7}
	n3.Left = n9
	n3.Right = n20
	n20.Left = n15
	n20.Right = n7

	// 队列
	fmt.Println("levelOrder result is: ", levelOrder(n3))
	// 迭代
	fmt.Println("levelOrderII result is: ", levelOrderII(n3))
}

func Test_invertTree(t *testing.T) {
	n4 := &TreeNode{Val: 4}
	n2 := &TreeNode{Val: 2}
	n7 := &TreeNode{Val: 7}
	n1 := &TreeNode{Val: 1}
	n3 := &TreeNode{Val: 3}
	n6 := &TreeNode{Val: 6}
	n9 := &TreeNode{Val: 9}

	n4.Left = n2
	n4.Right = n7
	n2.Left = n1
	n2.Right = n3
	n7.Left = n6
	n7.Right = n9

	levelOrder1 := levelOrder(n4)
	invertTree(n4)
	levelOrder2 := levelOrder(n4)
	fmt.Println("invertTree result is: ", levelOrder1, levelOrder2)

}

func Test_isSymmetric(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n22 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n33 := &TreeNode{Val: 3}
	n44 := &TreeNode{Val: 4}

	n1.Left = n2
	n1.Right = n22
	n2.Left = n3
	n2.Right = n4
	n22.Left = n44
	n22.Right = n33
	fmt.Println("isSymmetric result is: ", isSymmetric(n1)) // true
}

func Test_maxDepth(t *testing.T) {

	n3 := &TreeNode{Val: 3}
	n9 := &TreeNode{Val: 9}
	n20 := &TreeNode{Val: 20}
	n15 := &TreeNode{Val: 15}
	n7 := &TreeNode{Val: 7}

	n3.Left = n9
	n3.Right = n20
	n20.Left = n15
	n20.Right = n7

	fmt.Println("maxDepth result is: ", maxDepth(n3)) // 3

}

func Test_minDepth(t *testing.T) {

	n3 := &TreeNode{Val: 3}
	n9 := &TreeNode{Val: 9}
	n20 := &TreeNode{Val: 20}
	n15 := &TreeNode{Val: 15}
	n7 := &TreeNode{Val: 7}

	n3.Left = n9
	n3.Right = n20
	n20.Left = n15
	n20.Right = n7

	fmt.Println("minDepth result is: ", minDepth(n3)) // 2

}

func Test_countNodes(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{Val: 6}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6

	fmt.Println("countNode result is: ", countNodes(n1)) // 6
}

func Test_isBalance(t *testing.T) {
	n3 := &TreeNode{Val: 3}
	n9 := &TreeNode{Val: 9}
	n20 := &TreeNode{Val: 20}
	n15 := &TreeNode{Val: 15}
	n7 := &TreeNode{Val: 7}

	n3.Left = n9
	n3.Right = n20
	n20.Left = n15
	n20.Right = n7

	fmt.Println("isBalance result is: ", isBalanced(n3)) // true
}

func Test_isBalanceII(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n22 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n33 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n44 := &TreeNode{Val: 4}

	n1.Left = n2
	n1.Right = n22
	n2.Left = n3
	n3.Right = n33
	n3.Left = n4
	n3.Right = n44

	fmt.Println("isBalance result is: ", isBalanced(n1)) // false
}

func Test_binaryTreePaths(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n5 := &TreeNode{Val: 5}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n5
	fmt.Println("binaryTreePaths result is: ", binaryTreePaths(n1))
}

func Test_sumOfLeftLeaves(t *testing.T) {
	n3 := &TreeNode{Val: 3}
	n9 := &TreeNode{Val: 9}
	n20 := &TreeNode{Val: 20}
	n15 := &TreeNode{Val: 15}
	n7 := &TreeNode{Val: 7}

	n3.Left = n9
	n3.Right = n20
	n20.Left = n15
	n20.Right = n7

	fmt.Println("sumOfLeftLeaves result is: ", sumOfLeftLeaves(n3)) // 24
}

func Test_findBottomLeftValue(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{Val: 6}
	n7 := &TreeNode{Val: 7}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n3.Left = n5
	n3.Right = n6
	n5.Left = n7

	fmt.Println("findBottomLeftValue result is: ", findBottomLeftValue(n1)) // 7
}

func Test_hasPathSum(t *testing.T) {
	n5 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 4}
	n8 := &TreeNode{Val: 8}
	n11 := &TreeNode{Val: 11}
	n7 := &TreeNode{Val: 7}
	n2 := &TreeNode{Val: 2}
	n13 := &TreeNode{Val: 13}
	n44 := &TreeNode{Val: 4}
	n1 := &TreeNode{Val: 1}

	n5.Left = n4
	n5.Right = n8
	n4.Left = n11
	n11.Left = n7
	n11.Right = n2

	n8.Left = n13
	n8.Left = n44
	n44.Right = n1

	fmt.Println("hasPathSum result is: ", hasPathSum(n5, 22)) // true
}

func Test_buildTree(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	root := buildTree(inorder, postorder)
	fmt.Println("从中序与后序遍历序列构造二叉树，层序遍历结果：", levelOrder(root)) // [[3] [9 20] [15 7]]
}

func Test_buildTreeII(t *testing.T) {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}

	root := buildTreeII(preOrder, inOrder)
	fmt.Println("从前序与中序遍历序列构造二叉树，层序遍历结果：", levelOrder(root)) // [[3] [9 20] [15 7]]
}
