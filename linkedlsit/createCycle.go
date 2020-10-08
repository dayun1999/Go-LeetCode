package linkedlist

import "leetcode/definition"

//如何在链表的某个位置构建一个环
//步骤如下:
//1.遍历链表到第k个位置
//2.备份第k个链表
//3.继续遍历链表直至结束
//4.将最后一个结点和第k个结点连接起来
func MakeLoop(node *definition.SinglyLinkedListNode, k int) {
	current, count := node, 1
	for count < k {
		current = current.Next
		count++
	}
	jointPoint := current
	for current.Next != nil {
		current = current.Next
	}
	current.Next = jointPoint
}
