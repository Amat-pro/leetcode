package tree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preOrderTraversal 144 - 二叉树的前序遍历
func preOrderTraversal(root *TreeNode) []int {

	result := make([]int, 0, 10)
	// 中-左-右
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 中
		result = append(result, node.Val)
		// 左
		traversal(node.Left)
		// 右
		traversal(node.Right)
	}

	traversal(root)

	return result

}

// postOrderTraversal 145 - 二叉树的后序遍历  左 - 右 - 中
// inOrderTraversal 94 - 二叉树的中序遍历 左 - 中 - 右

// perOrderTraversalWithStack 二叉树前序-迭代遍历
func preOrderTraversalWithStack(root *TreeNode) []int {

	result := make([]int, 0, 10)
	// 中-左-右
	stack := make([]*TreeNode, 0, 10)
	stack = append(stack, root)
	for len(stack) > 0 {
		// pop
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 中 - 左 - 右
		result = append(result, n.Val)
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
		if n.Left != nil {
			stack = append(stack, n.Left) // 下次先pop左
		}

	}

	return result

}

// levelOrder - 102 二叉树的层序遍历 - 队列
func levelOrder(root *TreeNode) [][]int {

	result := make([][]int, 0, 10)

	// FIFO
	queue := make([]*TreeNode, 0, 10)

	queue = append(queue, root)
	for len(queue) > 0 {

		level := make([]int, 0, 10)
		size := len(queue)
		// FIFO - 已经在队列中的全部出队
		for i := 0; i < size; i++ {
			n := queue[0]
			queue = queue[1:]

			level = append(level, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}

		result = append(result, level)

	}
	return result

}

// levelOrder - 102 二叉树的层序遍历 - 递归
func levelOrderII(root *TreeNode) [][]int {

	result := make([][]int, 0, 10)

	var order func(node *TreeNode, depth int)
	order = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		index := depth
		if len(result) == index { // 如果当前depth没有存储位置，则为当前的depth创建存储位置
			result = append(result, []int{})
		}

		result[index] = append(result[index], node.Val)
		order(node.Left, depth+1)  // 递归左节点
		order(node.Right, depth+1) // 递归右节点

	}

	order(root, 0)

	return result

}

// invertTree 226 - 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {

	var invert func(node *TreeNode)
	invert = func(node *TreeNode) {

		if node == nil {
			return
		}

		node.Left, node.Right = node.Right, node.Left
		invert(node.Left)
		invert(node.Right)

	}

	invert(root)

	return root

}

// isSymmetric 101 - 对称二叉树
//
//	                               1
//	              2                                   2
//
//	     3                4                  4                   3
//	     .                。                  。                  .      // left.left==right.right && left.right==right.left
//	5         6       7        8        8         7          6        5
func isSymmetric(root *TreeNode) bool {

	var compare func(left *TreeNode, right *TreeNode) bool
	compare = func(left *TreeNode, right *TreeNode) bool {
		// 基本情况
		if left == nil && right == nil {
			return true
		}
		if (left == nil && right != nil) || (left != nil && right == nil) {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		leftEqual := compare(left.Left, right.Right)
		rightEqual := compare(left.Right, right.Left)

		return leftEqual && rightEqual

	}

	return compare(root.Left, root.Right)

}

// 前序求的就是深度，使用后序求的是高度
// 根节点的高度就是二叉树的最大深度

// maxDepth 104 - 二叉树的最大深度
func maxDepth(root *TreeNode) int {

	var order func(node *TreeNode) int
	order = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 后序遍历
		leftDepth := order(node.Left)
		rightDepth := order(node.Right)

		return 1 + max(leftDepth, rightDepth)

	}

	depth := order(root)
	return depth

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// minDepth - 111 二叉树的最小深度 （二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数，如果root.left!=nil, 最小深度不是1！！！）
func minDepth(root *TreeNode) int {

	var order func(node *TreeNode) int
	order = func(node *TreeNode) int {

		if node == nil {
			return 0
		}

		leftDepth := order(node.Left)
		rightDepth := order(node.Right)

		// node不是叶子结点的情况 需要找叶子结点！！
		if node.Left == nil && node.Right != nil {
			return 1 + rightDepth
		}
		if node.Left != nil && node.Right == nil {
			return 1 + leftDepth
		}

		//
		// node.Left == nil && node.Right == nil 的情况
		// node.Left != nil && node.Right != nil的情况
		return 1 + min(leftDepth, rightDepth)

	}

	return order(root)

}

// countNodes 222 - 完全二叉树的节点个数
func countNodes(root *TreeNode) int {

	var order func(node *TreeNode) int
	order = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftCount := order(node.Left)     // 左
		rightCount := order(node.Right)   // 右
		return 1 + leftCount + rightCount // 中
	}

	return order(root)

}

// isBalanced 110 - 平衡二叉树
func isBalanced(root *TreeNode) bool {

	var order func(node *TreeNode) (int, bool)
	order = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}

		leftHeight, leftBalance := order(node.Left)
		if !leftBalance {
			return -1, false
		}
		rightHeight, rightBalance := order(node.Right)
		if !rightBalance {
			return -1, false
		}

		if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
			return -1, false
		}

		return max(leftHeight, rightHeight) + 1, true

	}

	_, isBalance := order(root)

	return isBalance

}

// binaryTreePaths 257 - 二叉树的所有路径
func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0, 10)
	path := make([]int, 0, 10)

	var order func(node *TreeNode)
	order = func(node *TreeNode) {
		// 中 - 中的处理放在这里
		path = append(path, node.Val)

		// 终止条件 - 到达叶子结点
		if node.Left == nil && node.Right == nil {
			pathStr := []string{}
			for _, v := range path {
				pathStr = append(pathStr, strconv.Itoa(v))
			}
			result = append(result, strings.Join(pathStr, "->"))
			return
		}

		// 左
		if node.Left != nil {
			order(node.Left)
			// 回溯 - 左
			path = path[:len(path)-1]
		}
		// 右
		if node.Right != nil {
			order(node.Right)
			// 回溯 - 右
			path = path[:len(path)-1]
		}

	}

	order(root)

	return result

}

func sumOfLeftLeaves(root *TreeNode) int {

	var order func(node *TreeNode) int
	order = func(node *TreeNode) int {

		if node == nil {
			return 0
		}

		// 叶子结点
		if node.Left == nil && node.Right == nil {
			return 0
		}

		// 左
		leftValue := order(node.Left)

		// 左子树就是一个左叶子的情况
		if node.Left != node && node.Left.Left == nil && node.Left.Right == nil {
			leftValue = node.Left.Val
		}

		// 右
		rightValue := order(node.Right)

		// 中
		return leftValue + rightValue

	}

	return order(root)

}


// TODO 
// 找树左下角的值
// 路径总和
// 最大二叉树
// 合并两个二叉树

