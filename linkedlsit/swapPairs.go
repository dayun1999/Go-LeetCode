//--两两交换结点--https://leetcode-cn.com/problems/swap-nodes-in-pairs/

package linkedlist

//两两交换-递归
func SwapPairsRecursive(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	if head == nil || head.next == nil {
		return head
	}
	current := head.next
	head.next = SwapPairsRecursive(current.next)
	current.next = head
	return current
}

func SwapPairsPrevPointer(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	prev := &SinglyLinkedListNode{next: head}
	hint := prev
	for prev.next != nil && prev.next.next != nil {
		prev.next.next.next, prev.next.next, prev.next, prev = prev.next, prev.next.next.next, prev.next.next, prev.next
	}
	return hint.next
}
