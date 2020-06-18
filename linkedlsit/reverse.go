package linkedlist

//递归实现链表反转打印，不交换链表结点
//func reversePrint(head *SinglyLinkedListNode) {
//	if head != nil {
//		reversePrint(head.next)
//		fmt.Printf("%d->",head.data)
//	}
//}

//反转整个链表
func reverse(node *SinglyLinkedListNode) *SinglyLinkedListNode {
	//定义三个指针，一个prev为空，一个current代表当前结点，一个nextNode代表下一个结点
	current := node
	var prev *SinglyLinkedListNode = nil
	var nextNode *SinglyLinkedListNode = nil
	for current != nil {
		//存储当前结点的下一个结点
		nextNode = current.next
		//改变当前结点的下一个结点
		current.next = prev
		//把prev和current向下移动一位
		prev = current
		current = nextNode
	}
	//node = prev
	return prev
}
