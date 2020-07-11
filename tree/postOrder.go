package tree

import (
	"fmt"
)

//postOrder 后遍历
func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Printf("%d", root.Val)
}
