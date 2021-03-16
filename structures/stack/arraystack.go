package stack

import (
	"bytes"
	"fmt"
)

// 线性表版本实现 栈
type ArrayStack struct {
	data []interface{}
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{make([]interface{}, 0)}
}

func (stack *ArrayStack) GetSize() int {
	return len(stack.data)
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.GetSize() == 0
}

func (stack *ArrayStack) Push(e interface{}) {
	stack.data = append(stack.data, e)
}

func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	idx := len(stack.data)-1
	e := stack.data[idx]
	stack.data = stack.data[:idx]
	return e
}

func (stack *ArrayStack) Peek() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.data[len(stack.data)-1]
}

func (stack *ArrayStack) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Stack: ")
	buffer.WriteString("[")
	for i := 0; i < stack.GetSize(); i++ {
		buffer.WriteString(fmt.Sprint(stack.data[i]))
		if i != stack.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] top")

	return buffer.String()
}
