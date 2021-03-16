package stack

import (
	"github.com/Albert-Zuo/Data-Structures-And-Algorithm/structures/array"
	"github.com/Albert-Zuo/Data-Structures-And-Algorithm/structures/linked"
)

// 链表版本实现 栈
type LinkedStack struct {
	list array.List
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{
		list: linked.NewLinkedList(),
	}
}

func (stack *LinkedStack) GetSize() int {
	return stack.list.Size()
}

func (stack *LinkedStack) IsEmpty() bool {
	return stack.GetSize() == 0
}

func (stack *LinkedStack) Push(e interface{}) {
	stack.list.Append(e)
}

func (stack *LinkedStack) Pop() interface{} {
	return stack.list.Remove(stack.GetSize() - 1)
}

func (stack *LinkedStack) Peek() interface{} {
	e, _ := stack.list.Get(stack.GetSize() - 1)
	return e
}
