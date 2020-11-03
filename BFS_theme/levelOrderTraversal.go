package BFS_theme

import "leetcode/definition"

//题102: Binary Tree Level Order Traversal
//Given a binary tree, return the level order traversal of its
//nodes’ values. (ie, from left to right, level by level).
//Given binary tree [3,9,20,null,null,15,7],
//     3
//	/	 \
// 9     20
//     /    \
//    15     7
//return its level order traversal as:
//[
//[3],
//[9,20],
//[15,7]
//]

func levelOrder(root *definition.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*definition.TreeNode{}
	queue = append(queue, root)
	currentNum, nextLevelNum, res, tmp := 1, 0, [][]int{}, []int{}
	for len(queue) != 0 {
		if currentNum > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNum++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNum++
			}
			currentNum--
			tmp = append(tmp, node.Val)
			queue = queue[1:]
		}
		if currentNum == 0 {
			res = append(res, tmp)
			currentNum = nextLevelNum
			nextLevelNum = 0
			tmp = []int{}
		}
	}
	return res
}

//var result = []int{}
//
//func levelOrderTraversal(root *definition.TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//	result = append(result,root.Val)
//	result = append(result,levelOrderTraversal(root.Left)...)
//	result = append(result,levelOrderTraversal(root.Right)...)
//	return result
//}
