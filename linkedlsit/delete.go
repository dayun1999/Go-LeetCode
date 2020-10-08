//Delete a Node--https://www.hackerrank.com/challenges/delete-a-node-from-a-linked-list/problem
//made by code4EE
package linkedlist

import "leetcode/definition"

//链表的删除操作
//指定位置删除--01 常规操作
func Delete(head *definition.SinglyLinkedListNode, position int) *definition.SinglyLinkedListNode {
	if head == nil {
		return nil
	}
	if position == 0 {
		return head.Next
	}
	current := head
	var i int = 1
	for current.Next != nil {
		if i == position {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}
	return head
}

//指定位置删除--02 递归操作
func DeleteRecursive(head *definition.SinglyLinkedListNode, position int) *definition.SinglyLinkedListNode {
	if position == 0 {
		return head.Next
	}
	head.Next = DeleteRecursive(head.Next, position-1)
	return head
}
