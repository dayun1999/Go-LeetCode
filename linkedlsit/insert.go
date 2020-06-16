package linkedlist

//链表的插入操作
//一次性插入多个结点
func (list *SinglyLinkedList) insertMany(data ...int) {
	for _, item := range data {
		list.insertLast(item)
	}
}

//头插法
func (list *SinglyLinkedList) insertFirst(data int) {
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

//尾插法
func (list *SinglyLinkedList) insertLast(data int) {
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
