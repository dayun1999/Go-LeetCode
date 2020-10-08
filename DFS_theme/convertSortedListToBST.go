package DFS_theme

//题108：Convert Sorted Array to Binary Search Tree
//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。本题中，
//一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1

import "leetcode/definition"

func sortedArrayToBST(nums []int) *definition.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &definition.TreeNode{
		Val:   nums[len(nums)/2],
		Right: sortedArrayToBST(nums[len(nums)/2+1:]),
		Left:  sortedArrayToBST(nums[:len(nums)/2]),
	}
}

//自己写的
//func sortedListToBST(sortedList []int) *definition.TreeNode {
//	//根节点
//	root := &definition.TreeNode{
//		Val:   sortedList[len(sortedList)/2],
//		Right: nil,
//		Left:  nil,
//	}
//	for i:=len(sortedList)/2; i > 0 ; i-- {
//		convertArrayIntoBST(root.Left,sortedList[:len(sortedList)/2])
//		convertArrayIntoBST(root.Right,sortedList[len(sortedList)/2:])
//	}
//	return root
//}
//
//func convertArrayIntoBST(root *definition.TreeNode, nums []int) *definition.TreeNode {
//	if root == nil {
//		return root
//	}
//	for i := 0; i < len(nums); i++ {
//		if nums[i] < root.Val {
//			//放在左边
//			root.Left = &definition.TreeNode{
//				Val: nums[i],
//			}
//		} else if nums[i] > root.Val {
//			//放在右边
//			root.Right = &definition.TreeNode{
//				Val: nums[i],
//			}
//		}
//	}
//	return root
//}
