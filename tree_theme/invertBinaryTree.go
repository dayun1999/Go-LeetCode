package tree_theme

import "leetcode/definition"

//经典的“反转二叉树”
func invertTree(root *definition.TreeNode) *definition.TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Right, root.Left = root.Left, root.Right
	return root
}
