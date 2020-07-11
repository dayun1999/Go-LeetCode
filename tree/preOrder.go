package tree

import (
	"fmt"
)

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d", root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}
