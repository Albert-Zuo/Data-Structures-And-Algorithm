package linked

import (
	"errors"
	"fmt"
	"github.com/Albert-Zuo/Data-Structures-And-Algorithm/structures/array"
)

type Node struct {
	e    interface{}
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.e)
}

// LinkedList
type List struct {
	head *Node
	size int
}

func (list *List) Size() int {
	return list.size
}

func NewLinkedList() *List {
	return &List{
		head: nil,
		size: 0,
	}
}

// 返回链表是否为空
func (list *List) IsEmpty() bool {
	return list.size == 0
}

func (list *List) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.size {
		return nil, errors.New(fmt.Sprintf("index is out of range, required range: [0, %d) but got %d", list.size, index))
	}
	cur := list.head
	for i := 1; i <= index; i++ {
		cur = cur.next
	}
	return cur.e, nil

}

func (list *List) Set(index int, e interface{}) {
	if index < 0 || index >= list.size {
		panic("Set failed. Illegal index...")
	}
	cur := list.head
	for i := 1; i <= index; i++ {
		cur = cur.next
	}
	cur.e = e
}

func (list *List) Add(index int, e interface{}) {
	if index < 0 || index > list.size {
		panic("add failed, index is out of range")
	}

	if index == 0 {
		list.AddFirst(e)
	} else {
		// 获得待插入节点的前一个节点
		prev := list.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}

		// 插入新节点
		prev.next = &Node{e, prev.next}
		list.size++
	}
}

// 在链表头添加新的元素e
func (list *List) AddFirst(e interface{}) {

	list.head = &Node{e, nil}
	list.size++
}

func (list *List) Append(e interface{}) {
	list.Add(list.size, e)
}

func (list *List) Remove(index int) interface{} {
	if index < 0 || index >= list.size {
		panic(fmt.Sprintf(
			"remove failed, index is out of range, "+
				"required range: [0, %d] but got %d",
			list.size, index))
	}

	cur := list.head
	if index == 0 {
		list.head = list.head.next
		list.size--
		return cur.e
	}
	for i := 1; i < index; i++ {
		cur = cur.next
	}
	e := cur.next.e
	cur.next = cur.next.next
	list.size--
	return e
}

// 从链表中删除第一个元素，返回删除的元素
func (list *List) RemoveFirst() {
	list.Remove(0)
}

// 从链表中删除最后一个元素，返回删除的元素
func (list *List) RemoveLast() {
	list.Remove(list.size - 1)
}

func (list *List) Clean() {
	list.head = nil
	list.size = 0
}

func (list *List) Iterator() array.Iterator {
	return array.NewListIterator(list)
}
