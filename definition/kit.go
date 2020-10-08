package definition

//单链表的定义
type SinglyLinkedListNode struct {
	Data int
	Next *SinglyLinkedListNode
}

//链表的定义，包括头结点和尾结点
type SinglyLinkedList struct {
	Head *SinglyLinkedListNode
	Tail *SinglyLinkedListNode
}

//树的定义
type TreeNode struct {
	Val         int
	Right, Left *TreeNode
}
