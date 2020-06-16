//Delete a Node--https://www.hackerrank.com/challenges/delete-a-node-from-a-linked-list/problem
//made by code4EE
package linkedlist

//链表的删除操作
//指定位置删除--01 常规操作
func Delete(head *SinglyLinkedListNode, position int) *SinglyLinkedListNode {
	if head == nil {
		return nil
	}
	if position == 0 {
		return head.next
	}
	current := head
	var i int = 1
	for current.next != nil {
		if i == position {
			current.next = current.next.next
			break
		}
		current = current.next
	}
	return head
}

//指定位置删除--02 递归操作
func DeleteRecursive(head *SinglyLinkedListNode, position int) *SinglyLinkedListNode {
	if position == 0 {
		return head.next
	}
	head.next = DeleteRecursive(head.next, position-1)
	return head
}
