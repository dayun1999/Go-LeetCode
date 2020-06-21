//合并两个有序链表--https://leetcode-cn.com/problems/merge-two-sorted-lists/
//made by code4EE

package linkedlist

//递归
func MergeTwoListsRecursive(l1 *SinglyLinkedListNode, l2 *SinglyLinkedListNode) *SinglyLinkedListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.data < l2.data {
		l1.next = MergeTwoListsRecursive(l1.next, l2)
		return l1
	} else {
		l2.next = MergeTwoListsRecursive(l1, l2.next)
		return l2
	}

}

//迭代
func MergeTwoLists(l1 *SinglyLinkedListNode, l2 *SinglyLinkedListNode) *SinglyLinkedListNode {
	newNode := &SinglyLinkedListNode{}
	prev := newNode
	for l1 != nil && l2 != nil {
		if l1.data < l2.data {
			newNode.next = l1
			l1 = l1.next
		} else {
			newNode.next = l2
			l2 = l2.next
		}
		newNode = newNode.next
	}
	if l1 != nil {
		newNode.next = l1
	}
	if l2 != nil {
		newNode.next = l2
	}
	return prev.next
}
