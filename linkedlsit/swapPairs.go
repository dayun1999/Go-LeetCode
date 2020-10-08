//--两两交换结点--https://leetcode-cn.com/problems/swap-nodes-in-pairs/

package linkedlist

import "leetcode/definition"

//两两交换-递归
func SwapPairsRecursive(head *definition.SinglyLinkedListNode) *definition.SinglyLinkedListNode {
	if head == nil || head.Next == nil {
		return head
	}
	current := head.Next
	head.Next = SwapPairsRecursive(current.Next)
	current.Next = head
	return current
}

//func SwapPairsPrevPointer(head *SinglyLinkedListNode) *SinglyLinkedListNode {
//	prev := &SinglyLinkedListNode{Next: head}
//	hint := prev
//	for prev.Next != nil && prev.Next.Next != nil {
//		prev.Next.Next.Next, prev.Next.Next, prev.Next, prev = prev.Next, prev.Next.Next.Next, prev.Next.Next, prev.Next
//	}
//	return hint.Next
//}
