package tree_theme

import (
	"fmt"
	"leetcode/definition"
)

//层次遍历
func LevelOrder(root *definition.TreeNode) {
	var level = 1
	for printLevelOrder(root, level) {
		level++
	}
}

func printLevelOrder(root *definition.TreeNode, level int) bool {
	if root == nil {
		return false
	}
	if level == 1 {
		fmt.Printf("%d ", root.Val)
		return true
	}
	left := printLevelOrder(root.Left, level-1)
	right := printLevelOrder(root.Right, level-1)
	return left || right
}
