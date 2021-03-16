package stack

import (
	"testing"
)

var tests = []string{
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
}

func TestArrayStack_Push(t *testing.T) {
	stack := NewArrayStack()
	for _, v := range tests {
		stack.Push(v)
	}

	if stack.GetSize() != len(tests) {
		t.Error("ArrayStack_Push error")
	}

}

func TestLinkedStack_Push(t *testing.T) {
	stack := NewLinkedStack()
	for _, v := range tests {
		stack.Push(v)
	}

	if stack.GetSize() != len(tests) {
		t.Error("ArrayStack_Push error")
	}
}

func TestArrayStack_Pop(t *testing.T) {
	stack := NewArrayStack()
	for _, v := range tests {
		stack.Push(v)
	}

	for i := range tests {
		e := stack.Pop()
		if e != tests[len(tests)-1-i] {
			t.Errorf("ArrayStack_Pop error, current value is %s , but got %s ", tests[len(tests)-1-i], e)
		}
	}

}

func TestLinkedStack_Pop(t *testing.T) {

	stack := NewLinkedStack()
	for _, v := range tests {
		stack.Push(v)
	}

	for i := range tests {
		e := stack.Pop()
		if e != tests[len(tests)-1-i] {
			t.Errorf("LinkedStack_Pop error, current value is %s , but got %s ", tests[len(tests)-1-i], e)
		}
	}
}

func TestArrayStack_Peek(t *testing.T) {
	stack := NewArrayStack()
	for _, v := range tests {
		stack.Push(v)
		e := stack.Peek()
		if e != v {
			t.Errorf("ArrayStack_Peek error, current value is %s , but got %s ", v, e)
		}
	}

}

func TestLinkedStack_Peek(t *testing.T) {

	stack := NewLinkedStack()
	for _, v := range tests {
		stack.Push(v)
		e := stack.Peek()
		if e != v {
			t.Errorf("LinkedStack_Peek error, current value is %s , but got %s ", v, e)
		}
	}
}
