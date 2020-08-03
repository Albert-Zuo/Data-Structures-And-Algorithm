package array

import (
	"bytes"
	"errors"
	"fmt"
)

type List interface {
	Size() int                                // 获取 List 长度

	Get(index int) (interface{}, error)       // 根据索引获取 element

	Set(index int, e interface{})             // 修改索引对应的值

	Add(index int, e interface{})             // 在索引位置添加新的 element

	Append(e interface{})                     // 末尾添加 element

	Remove(index int) interface{}             // 根据索引删除 element

	Clean()                                   // 清理数组内存

	Iterator() Iterator                       // 获取迭代器
}

type Array struct {
	size int
	data []interface{}
}

// 自定义 capacity
func NewArrayByCapacity(capacity int) *Array {
	if capacity < 0 {
		panic(
			fmt.Sprintf("Invalid parameter, capacity no negative number, got capacity = %d", capacity))
	}
	return &Array{
		data: make([]interface{}, capacity),
		size: 0,
	}
}

// 默认构造函数 容量=10
func NewArray() *Array {
	list := new(Array)

	list.data = make([]interface{}, 10)
	list.size = 0
	return list
}

func (list *Array) Clean() {
	list.data = make([]interface{}, 10)
	list.size = 0
}

func (list *Array) Size() int {
	return list.size
}

func (list *Array) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.size {
		return nil, errors.New("index is out of range")
	}
	return list.data[index], nil
}

func (list *Array) Set(index int, e interface{}) {
	list.checkIndex(index)
	list.data[index] = e
}

func (list *Array) Add(index int, e interface{}) {

	if index < 0 || index > list.size {
		panic(fmt.Sprintf("index is out of range, required range: [0, %d] but got %d", list.size, index))
	}

	if list.size == len(list.data) {
		list.resize(2 * list.size)
	}

	for i := list.size - 1; i >= index ; i++ {
		list.data[i + 1] = list.data[i]
	}
	list.data[index] = e
	list.size++
}

// 向所有元素后添加一个新元素
func (list *Array) Append(e interface{}) {
	list.data[list.size] = e
	list.size++
}

// 从数组中删除index位置的元素, 返回删除的元素
func (list *Array) Remove(index int) interface{} {

	list.checkIndex(index)
	result := list.data[index]

	// 删除index位置的元素
	list.data = append(list.data[:index], list.data[index+1:]...)
	list.size--

	// resize array ...
	currLength := len(list.data)
	if list.size == currLength/4 && currLength/2 != 0 {
		list.resize(currLength / 2);
	}
	return result
}

// 扩容缩容
func (list *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)

	for _, v := range list.data {
		newData = append(newData, v)
	}
	list.data = newData
}

// toString
func (list *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", list.size, len(list.data)))
	buffer.WriteString("[")
	for i := 0; i < list.size; i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(list.data[i]))
		if i != (list.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}

// 检查索引是否生效
func (list *Array) checkIndex(index int) {
	if index < 0 || index >= list.size {
		panic(fmt.Sprintf("index is out of range, required range: [0, %d) but got %d", list.size, index))
	}
}

// 获取迭代器
func (list *Array) Iterator() Iterator {
	iterator := &ListIterator{
		currIndex: 0,
		list:      list,
	}
	return iterator
}
