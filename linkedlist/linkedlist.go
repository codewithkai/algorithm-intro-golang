package linkedlist

import (
	"strconv"
	"strings"
)

type ListNode struct {
	val  int
	next *ListNode
}

type LinkedList struct {
	head ListNode //注意这里不是指针，而是一个节点，即所谓的头结点，本身不存储数据，可以简化链表处理
	// len int //还可以记录链表长度，插入删除时更新即可
}

func NewLinkedList() *LinkedList {
	return &LinkedList{} //初始化了头结点
}

// 头部添加元素
func (list *LinkedList) AddAtHead(val int) {
	newNode := &ListNode{val: val, next: list.head.next}
	list.head.next = newNode
}

//尾部添加元素
func (list *LinkedList) AddAtTail(val int) {
	current := &list.head
	for current.next != nil { //循环完成之后current指向最后一个节点，而不是nil，不然没法操作了
		current = current.next
	}
	newNode := &ListNode{val: val}
	current.next = newNode
}

// Reverse 翻转链表
func (list *LinkedList) Reverse() {
	var prev *ListNode //利用默认值nil
	current := list.head.next
	for current != nil {
		// 这种语法很方便，省去了中间变量。比如Java就不可以这样，而Python可以
		// 翻转操作就一步： current.next = prev
		// prev, current = current, current.next 往下移动两个工作指针，一个是prev, 一个current
		prev, current, current.next = current, current.next, prev
	}
	list.head.next = prev
}

// 根据val 删除节点
func (list *LinkedList) Delete(val int) {
	prev := &list.head
	current := list.head.next
	for current != nil {
		if current.val == val {
			// 如果没有头节点，这里就没这么简洁了，需要做额外的判断
			prev.next = current.next
			return
		}
		prev, current = current, current.next
	}
}

// 打印链表
func (list *LinkedList) Print() string {
	current := list.head.next
	var vals []string
	for current != nil {
		vals = append(vals, strconv.Itoa(current.val))
		current = current.next
	}
	vals = append(vals, "nil")
	return strings.Join(vals, "->")
}
