//--两两交换结点--https://leetcode-cn.com/problems/swap-nodes-in-pairs/

package recursive

type ListNode struct {
	data int
	next *ListNode
}

//两两交换
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return nil
	}
	head, head.next = head.next, head
	head = head.next.next
	return swapPairs(head)
}
