//LCA- Lowest Common Ancestor 最小共同祖先
package tree

//二叉搜索树中为两个结点寻找最小共同祖先
func LCA(root *TreeNode, target1 int, target2 int) *TreeNode {
	//利用深度优先方法
	if root == nil {
		return nil
	}
	if root.Val < target1 && root.Val < target2 {
		//当该根节点都小于两个目标结点(的值)，递归找寻右子树
		return LCA(root.Right, target1, target2)
	} else if root.Val > target1 && root.Val > target2 {
		return LCA(root.Left, target1, target2)
	} else {
		return root
	}
}
