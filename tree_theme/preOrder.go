package tree_theme

import (
	"fmt"
	"leetcode/definition"
)

func PreOrder(root *definition.TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d", root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}
