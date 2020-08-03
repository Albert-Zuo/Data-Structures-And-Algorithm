package array

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

func TestArray_Add(t *testing.T) {
	list := NewArray()
	// 填充数据
	for i, v := range tests {
		list.Add(i, v)
	}
	if list.Size() != len(tests) {
		t.Errorf("Array lenght is %d but got %d ", len(tests), list.Size())
	}
}

func TestArray_Append(t *testing.T) {
	list := NewArray()
	// 填充数据
	for _, value := range tests {
		list.Append(value)
	}
	if list.Size() != len(tests) {
		t.Errorf("Array lenght is %d but got %d ", len(tests), list.Size())
	}
}

func TestArray_Iterator(t *testing.T) {

	list := NewArray()
	// 填充数据
	for _, value := range tests {
		list.Append(value)
	}
	fmt.Println(list)
	// 遍历删除一个元素
	for iterator := list.Iterator(); iterator.HasNext(); {
		next, _ := iterator.Next()
		if next == "e" {
			iterator.Remove()
		}
	}
	for iterator := list.Iterator(); iterator.HasNext(); {
		next, _ := iterator.Next()
		if next == "e" {
			t.Errorf("Array中的 %s 删除失败！", next)
		}
	}

}
