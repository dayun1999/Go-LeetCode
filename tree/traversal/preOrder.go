package traversal

import (
	"fmt"
	"leetcode/tree"
)

func PreOrder(root *tree.TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d", root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}
