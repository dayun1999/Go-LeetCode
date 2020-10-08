package DFS_theme

//题104: Maximum Depth of Binary Tree

import "leetcode/definition"

func maxDepth(root *definition.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	//left := maxDepth(root.Left)
	//right := maxDepth(root.Right)
	//if left < right {
	//	return right
	//}
	//return left
}

//返回最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
