package tree_theme

import (
	"fmt"
	"leetcode/definition"
)

//postOrder 后遍历
func PostOrder(root *definition.TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Printf("%d", root.Val)
}
