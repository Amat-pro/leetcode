package temp

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	// 前序遍历
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	// 中
	root1.Val = root1.Val + root2.Val
	// 左
	root1.Left = mergeTrees(root1.Left, root2.Right)
	root1.Right = mergeTrees(root1.Right, root2.Right)

	return root1

}