//二叉树的创建和插入操作
package tree_theme

import "leetcode/definition"

//建立二叉搜索树-Binary Search Tree
func NewBST(value ...int) *definition.TreeNode {
	var root *definition.TreeNode
	for _, v := range value {
		root = insert(root, v)
	}
	return root
}

//二叉树插入结点的规则:比根节点小插在左边，和根节点相等本方法选择忽略，比根节点大的插在右边，递归进行
func insert(root *definition.TreeNode, value int) *definition.TreeNode {
	if root == nil {
		return &definition.TreeNode{value, nil, nil}
	}
	if value > root.Val {
		root.Right = insert(root.Right, value)
	} else {
		root.Left = insert(root.Left, value)
	}
	return root
}
