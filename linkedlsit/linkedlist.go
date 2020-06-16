package linkedlist

//单链表的定义在这
//结点的定义
type SinglyLinkedListNode struct {
	data int
	next *SinglyLinkedListNode
}

//链表的定义，包括头结点和尾结点
type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}
