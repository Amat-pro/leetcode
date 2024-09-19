package tree

// preOrderTraversal 前序遍历
func preOrderTraversal(root *TreeNode) []int {

	result := []int{}
	var traversalFn func(node *TreeNode)

	traversalFn = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 左 - 中 - 右
		// 前序遍历：中左右
		result = append(result, node.Val)
		traversalFn(node.Left)
		traversalFn(node.Right)

	}

	traversalFn(root)

	return result
}

// preOrderTraversal_V2 层序遍历
func preOrderTraversal_V2(root *TreeNode) []int {

	result := []int{}
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) > 0 {
		// 前序遍历：中 - 左 - 右
		node := stack[len(stack)-1]  
		stack = stack[0:len(stack)-1]

		result = append(result, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

	}

	return result


}

func inOrderTraversal(root *TreeNode) []int {
	result := []int{}
	var traversalFn func(node *TreeNode)

	traversalFn = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 左 - 中 - 右
		// 中序遍历：左中右
		traversalFn(node.Left)
		result = append(result, node.Val)
		traversalFn(node.Right)

	}

	traversalFn(root)

	return result
}

func postOrderTraversal(root *TreeNode) []int {
	result := []int{}
	var traversalFn func(node *TreeNode)

	traversalFn = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 左 - 中 - 右
		// 后序遍历：左右中
		traversalFn(node.Left)
		traversalFn(node.Right)
		result = append(result, node.Val)

	}

	traversalFn(root)

	return result
}