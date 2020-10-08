package DFS_theme

import "leetcode/definition"

//题目描述
//Given two binary trees, write a function to check if they are the same or not.
//
//Two binary trees are considered the same if they are structurally identical and the nodes have the same value.
//

func isSameTree(p *definition.TreeNode, q *definition.TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}

}

//自己写的
//var res []int
//
//func checkSameTree(root1 , root2 *definition.TreeNode) bool {
//	res = []int{}
//	nodes := dfs(root1)
//	nodes2 := dfs(root2)
//	flag := true
//	for i := 0; i < len(nodes); i++ {
//		if nodes[i] != nodes2[i] {
//			flag = false
//		}
//	}
//	return flag
//}
//
//func dfs(root *definition.TreeNode) []int {
//	if root == nil {
//		return res
//	}
//	res = append(res, dfs(root.Right)...)
//	res = append(res, dfs(root.Left)...)
//	return res
//}
