package tree_theme

import (
	"fmt"
	"leetcode/definition"
)

//inOrder 中序遍历
func InOrder(root *definition.TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Printf("%d", root.Val)
	InOrder(root.Right)
}
