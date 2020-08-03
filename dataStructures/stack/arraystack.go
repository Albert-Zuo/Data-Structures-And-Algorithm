package stack

import (
	"bytes"
	"fmt"
	"github.com/Albert-Zuo/Data-Structures-And-Algorithm/dataStructures/array"
)

// 线性表版本实现 栈
type ArrayStack struct {
	list array.List
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		list: array.NewArray(),
	}
}

func (stack *ArrayStack) GetSize() int {
	return stack.list.Size()
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.GetSize() == 0
}

func (stack *ArrayStack) Push(e interface{}) {
	stack.list.Append(e)
}

func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.list.Remove(stack.GetSize() - 1)
}

func (stack *ArrayStack) Peek() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	e, _ := stack.list.Get(stack.GetSize() - 1)
	return e

}

func (stack *ArrayStack) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Stack: ")
	buffer.WriteString("[")
	for i := 0; i < stack.GetSize(); i++ {
		buffer.WriteString(fmt.Sprint(stack.list.Get(i)))
		if i != stack.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] top")

	return buffer.String()
}
