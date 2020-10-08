package linkedlist

import "leetcode/definition"

type linkedlist definition.SinglyLinkedList

//链表的插入操作
//头插法
func (list *linkedlist) InsertFirst(items ...int) {
	for _, data := range items {
		node := &definition.SinglyLinkedListNode{
			Data: data,
			Next: nil,
		}
		if list.Head == nil {
			list.Head = node
		} else {
			node.Next = list.Head
		}
		list.Tail = node
	}
}

//尾插法
func (list *linkedlist) InsertLast(items ...int) {
	for _, data := range items {
		node := &definition.SinglyLinkedListNode{
			Data: data,
			Next: nil,
		}
		if list.Head == nil {
			list.Head = node
		} else {
			list.Tail.Next = node
		}
		list.Tail = node
	}
}
