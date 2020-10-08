package DFS_theme

import "leetcode/definition"

func recoverTree(root *definition.TreeNode) {
	var prev, target1, target2 *definition.TreeNode
	_, target1, target2 = inOrderTraverse(root, prev, target1, target2)
	if target1 != nil && target2 != nil {
		target1.Val, target2.Val = target2.Val, target1.Val
	}
}

//先序遍历，找到有问题的两个节点
func inOrderTraverse(root, prev, target1, target2 *definition.TreeNode) (*definition.TreeNode, *definition.TreeNode, *definition.TreeNode) {
	if root == nil {
		return prev, target1, target2
	}
	prev, target1, target2 = inOrderTraverse(root.Left, prev, target1, target2)
	if prev != nil && prev.Val > root.Val {
		if target1 == nil {
			target1 = prev
		}
		target2 = root
	}
	prev = root
	prev, target1, target2 = inOrderTraverse(root.Right, prev, target1, target2)

	return prev, target1, target2
}
