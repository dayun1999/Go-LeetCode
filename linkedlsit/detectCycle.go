package linkedlist

import "leetcode/definition"

//检查链表里面是否有环

//使用快慢指针(Floyd's Algorithm)
func DetectCycleByPointer(node *definition.SinglyLinkedListNode) bool {
	if node == nil || node.Next == nil {
		return false
	}
	slow, fast := node, node.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

//hashing方法
func DetectCycleByHashing(node *definition.SinglyLinkedListNode) bool {
	if node == nil {
		return false
	}
	set := make(map[*definition.SinglyLinkedListNode]int)
	current := node
	for current.Next != nil {
		set[current]++
		if num, ok := set[current]; ok && num == 2 { //wrong!!!
			return true
		}
		current = current.Next
	}
	return false
}
