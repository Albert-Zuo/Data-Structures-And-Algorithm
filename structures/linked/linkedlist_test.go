package linked

import (
	"fmt"
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

func TestLinkedList_Add(t *testing.T) {

	list := NewLinkedList()
	for i, v := range tests {
		list.Add(i, v)
	}

	for i, v := range tests {
		e, _ := list.Get(i)
		if e != v {
			t.Errorf("current value is %s , but got %s ", v, e)
		}
	}
}

func TestLinkedList_Set(t *testing.T) {

	list := NewLinkedList()
	for i, v := range tests {
		list.Add(i, v)
	}
	value := "g"
	for i, v := range tests {
		// 更改值
		list.Set(i, value)

		e, _ := list.Get(i)
		if e == v {
			t.Errorf("current value is %s , but got %s ", value, e)
		}
	}
}

func TestLinkedList_Iterator(t *testing.T) {
	list := NewLinkedList()
	for i, v := range tests {
		list.Add(i, v)
	}

	for iterator := list.Iterator(); iterator.HasNext(); {

		index := iterator.GetIndex()
		e, _ := list.Get(index)
		next, _ := iterator.Next()
		if e != next {
			t.Errorf("current value is %s , but got %s ", e, next)
		}

	}
}

func TestLinkedList_Remove(t *testing.T) {

	list := NewLinkedList()
	for i, v := range tests {
		list.Add(i, v)
	}

	for range tests {
		e := list.Remove(list.Size() - 1)
		fmt.Println(e)
	}

	if list.Size() != 0 {
		t.Error("remove failed !")
	}

}
