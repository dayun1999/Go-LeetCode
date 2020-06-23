package traversal

import (
	"fmt"
	"leetcode/tree"
)

//inOrder 中序遍历
func InOrder(root *tree.TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Printf("%d", root.Val)
	InOrder(root.Right)
}
