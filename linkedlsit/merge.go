//合并两个有序链表--https://leetcode-cn.com/problems/merge-two-sorted-lists/
//made by code4EE

package linkedlist

import "leetcode/definition"

//递归
func MergeTwoListsRecursive(l1 *definition.SinglyLinkedListNode, l2 *definition.SinglyLinkedListNode) *definition.SinglyLinkedListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Data < l2.Data {
		l1.Next = MergeTwoListsRecursive(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoListsRecursive(l1, l2.Next)
		return l2
	}

}

//迭代
func MergeTwoLists(l1 *definition.SinglyLinkedListNode, l2 *definition.SinglyLinkedListNode) *definition.SinglyLinkedListNode {
	newNode := &definition.SinglyLinkedListNode{}
	prev := newNode
	for l1 != nil && l2 != nil {
		if l1.Data < l2.Data {
			newNode.Next = l1
			l1 = l1.Next
		} else {
			newNode.Next = l2
			l2 = l2.Next
		}
		newNode = newNode.Next
	}
	if l1 != nil {
		newNode.Next = l1
	}
	if l2 != nil {
		newNode.Next = l2
	}
	return prev.Next
}
