package tree

import (
	"fmt"
)

//inOrder 中序遍历
func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Printf("%d", root.Val)
	InOrder(root.Right)
}
