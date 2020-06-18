package linkedlist

//检查链表里面是否有环

//使用快慢指针(Floyd's Algorithm)
func DetectCycleByPointer(node *SinglyLinkedListNode) bool {
	if node == nil || node.next == nil {
		return false
	}
	slow, fast := node, node.next
	for slow != fast {
		if fast == nil || fast.next == nil {
			return false
		}
		fast = fast.next.next
		slow = slow.next
	}
	return true
}

//hashing方法
func DetectCycleByHashing(node *SinglyLinkedListNode) bool {
	if node == nil {
		return false
	}
	set := make(map[*SinglyLinkedListNode]int)
	current := node
	for current.next != nil {
		set[current]++
		if num, ok := set[current]; ok && num == 2 { //wrong!!!
			return true
		}
		current = current.next
	}
	return false
}
