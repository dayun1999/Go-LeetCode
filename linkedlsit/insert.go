package linkedlist

//链表的插入操作

//头插法
func (list *SinglyLinkedList) InsertFirst(items ...int) {
	for _, data := range items {
		node := &SinglyLinkedListNode{
			data: data,
			next: nil,
		}
		if list.head == nil {
			list.head = node
		} else {
			node.next = list.head
		}
		list.tail = node
	}
}

//尾插法
func (list *SinglyLinkedList) InsertLast(items ...int) {
	for _, data := range items {
		node := &SinglyLinkedListNode{
			data: data,
			next: nil,
		}
		if list.head == nil {
			list.head = node
		} else {
			list.tail.next = node
		}
		list.tail = node
	}
}
