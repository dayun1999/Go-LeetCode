package DFS_theme

import "leetcode/definition"

//题目描述
//Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
//将原来的树进行反转之后，再判断反转之后的树和原来的是否相等
func isSymmetricTree(root *definition.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(invertTree(root.Left), root.Right)
}

//反转二叉树的每一个结点
func invertTree(root *definition.TreeNode) *definition.TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

//判断两个树是否相等
//func isSameTree(p , q *definition.TreeNode) bool {
//	if p == nil && q == nil {
//		return true
//	} else if p != nil && q != nil {
//		return isSameTree(p.Left,q.Left) && isSameTree(p.Right, q.Right)
//	} else {
//		return false
//	}
//}
