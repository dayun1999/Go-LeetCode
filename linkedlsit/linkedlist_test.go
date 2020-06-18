package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	var list = SinglyLinkedList{}
	//插入
	list.InsertLast(9, 2, 3, 45, 10, 12)
	t.Logf("原始链表如下:%s\n", ToString(list.head))

	//删除某个位置的结点
	DeleteRecursive(list.head, 1)
	t.Logf("删除结点之后的链表如下:%s\n", ToString(list.head))
	//两两交换
	list.head = SwapPairsPrevPointer(list.head)
	t.Logf("两两交换之后的链表如下:%s\n", ToString(list.head))
	//反转打印链表
	t.Logf("链表反转之后如下：%s\n", ToString(reverse(list.head)))

}

func TestDetectingLoop(t *testing.T) {
	var list1 = SinglyLinkedList{}
	//插入
	list1.InsertLast(9, 2, 3, 45, 10, 12)
	t.Logf("原始链表如下:%s\n", ToString(list1.head))
	//创建一个环
	//MakeLoop(list1.head,1)
	t.Logf("是否有环: %v\n", DetectCycleByPointer(list1.head))
	//t.Logf("加环之后的链表如下: %s\n",ToString(list1.head))
	t.Logf("是否有环: %v\n", DetectCycleByHashing(list1.head))
}
