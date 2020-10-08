package DFS_theme

//题98:Validate Binary Search Tree
//Given a binary tree, determine if it is a valid binary search tree (BST).
//
//Assume a BST is defined as follows:
//
//The left subtree of a node contains only nodes with keys less than the node’s key.
//The right subtree of a node contains only nodes with keys greater than the node’s key.
//Both the left and right subtrees must also be binary search trees.

import (
	"leetcode/definition"
	"math"
)

//法二
//直接按照定义，左节点 < 根节点 < 右节点
func isValidBST2(root *definition.TreeNode) bool {
	return validate(root, math.Inf(-1), math.Inf(1))
}

func validate(root *definition.TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	return v < max && v > min && validate(root.Left, min, v) && validate(root.Right, v, max)
}

//法一
//将BST的节点按照左中右输出到数组，如果出现逆序就说明不是BST
func isVliadBST1(root *definition.TreeNode) bool {
	res := []int{}
	inOrder(root, &res)
	//遍历，判断是否有逆序的
	for i := 1; i < len(res); i++ {
		if res[i-1] >= res[i] {
			return false
		}
	}
	return true
}

func inOrder(root *definition.TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}
