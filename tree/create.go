//二叉树的创建和插入操作
package tree

//建立二叉搜索树-Binary Search Tree
func NewBST(value ...int) *TreeNode {
	var root *TreeNode
	for _, v := range value {
		root = insert(root, v)
	}
	return root
}

//二叉树插入结点的规则:比根节点小插在左边，和根节点相等本方法选择忽略，比根节点大的插在右边，递归进行
func insert(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{value,nil,nil}
	}
	if value > root.Val {
		root.Right = insert(root.Right, value)
	} else {
		root.Left = insert(root.Left, value)
	}
	return root
}
