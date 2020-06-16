package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := SinglyLinkedList{}
	list.insertMany(9, 2, 3, 45, 10, 12)
	t.Logf("原始链表如下:%s\n", toString(list.head))
	DeleteRecursive(list.head, 1)
	t.Logf("删除结点之后的链表如下:%s\n", toString(list.head))
}
