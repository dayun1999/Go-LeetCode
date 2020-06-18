package linkedlist

import (
	"fmt"
)

//打印链表
func ToString(head *SinglyLinkedListNode) string {
	if head == nil {
		return "the linkedlist is empty."
	}
	current := head
	s := ""
	for current != nil {
		s += fmt.Sprintf("%d->", current.data)
		current = current.next
	}
	return s
}
