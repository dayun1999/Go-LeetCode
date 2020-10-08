package linkedlist

import (
	"fmt"
	"leetcode/definition"
)

//打印链表
func ToString(head *definition.SinglyLinkedListNode) string {
	if head == nil {
		return "the linkedlist is empty."
	}
	current := head
	s := ""
	for current != nil {
		s += fmt.Sprintf("%d->", current.Data)
		current = current.Next
	}
	return s
}
