package linkedlist

import (
	"fmt"
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
	list.head = SwapPairsRecursive(list.head)
	t.Logf("两两交换之后的链表如下:%s\n", ToString(list.head))
	//反转打印链表
	t.Logf("链表反转之后如下：%s\n", ToString(reverse(list.head)))

}

//测试链表有没有环
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

//测试合并两个有序链表
func TestMergeTwoLists(t *testing.T) {
	var list2 = SinglyLinkedList{}
	var list3 = SinglyLinkedList{}
	//第一个链表--1,3,4,5
	list2.InsertLast(1, 3, 4, 5)
	//第二个链表--1,2,4,6,7,8
	list3.InsertLast(1, 2, 4, 6, 7, 8)
	//合并链表
	expected := "1->1->2->3->4->4->5->6->7->8->"
	newHead := MergeTwoLists(list2.head, list3.head)
	//t.Logf("合并链表如下: %s\n",ToString(newHead))
	actual := fmt.Sprintf("%v", ToString(newHead))
	if expected != actual {
		t.Errorf("期待得到的是%v,然而得到的却是%v\n", expected, actual)
	}
}
