package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	list := NewLinkedList()

	t.Log(list.Print())
	list.AddAtHead(1)
	t.Log(list.Print())
	list.AddAtTail(2)
	t.Log(list.Print())
	list.AddAtHead(0)
	t.Log(list.Print())

	list.Reverse()
	t.Log(list.Print())

	list.Delete(0)
	t.Log(list.Print())
	list.Delete(2)
	t.Log(list.Print())
	list.Delete(1)
	t.Log(list.Print())
}
