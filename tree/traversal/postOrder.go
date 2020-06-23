package traversal

import (
	"fmt"
	"leetcode/tree"
)

//postOrder 后遍历
func PostOrder(root *tree.TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Printf("%d", root.Val)
}
