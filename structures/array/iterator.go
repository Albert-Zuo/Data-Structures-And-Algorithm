package array

import "errors"

type Iterator interface {
	HasNext() bool              // 是否存在元素
	Next() (interface{}, error) // 获取元素
	Remove()                    // 删除元素
	GetIndex() int              // 获取元素索引
}

type Iterable interface {
	Iterator() Iterator
}

type ListIterator struct {
	list      List
	currIndex int // 当前指向的索引
}

func (iterator *ListIterator) HasNext() bool {
	return iterator.currIndex < iterator.list.Size()
}

func (iterator *ListIterator) Next() (interface{}, error) {
	if !iterator.HasNext() {
		return nil, errors.New("There is no next element in the current array...")
	}
	element, err := iterator.list.Get(iterator.currIndex)
	iterator.currIndex++
	return element, err
}

func (iterator *ListIterator) Remove() {
	iterator.currIndex--
	iterator.list.Remove(iterator.currIndex)
}

func (iterator *ListIterator) GetIndex() int {
	return iterator.currIndex
}

func NewListIterator(list List) *ListIterator {
	return &ListIterator{
		list:      list,
		currIndex: 0,
	}
}
