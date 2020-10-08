package tree_theme

import (
	"leetcode/definition"
	"testing"
)

var root *definition.TreeNode

func TestCreateTree(t *testing.T) {
	root = NewBST(7, 4, 5, 6, 3, 9, 10)
	LevelOrder(root)
	//		    7
	//	4				9
	//3		5				10
	//			6

}

//测试最小共同祖先 LCA
func TestLCA(t *testing.T) {
	node := LCA(root, 6, 10)
	expected := 7
	if node.Val != expected {
		t.Errorf("期望得到的是%d,然而得到的是%d\n", expected, node.Val)
	}
}
