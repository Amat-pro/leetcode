package temp

// invertTree 翻转二叉树 - 前序遍历
func invertTree(root *TreeNode) *TreeNode {

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root

}