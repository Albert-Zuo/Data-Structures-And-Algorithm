package sort

import (
	"reflect"
)

// 快排
type arrayQuick struct {
	Array
}

func NewArrayQuick() ArraySorter {
	return &arrayQuick{Array{Type: "quickSort"}}
}

func(a *arrayQuick) Sort(s []int) {
	a.slice = s
	arrayQuickSort(a.slice[:], 0, len(a.slice)-1)
}

func arrayQuickSort(s []int, l, r int) {
	if l >= r {
		return
	}
	key := s[l]
	i, j := l-1, r+1
	for i < j {
		for {
			i++
			if s[i] >= key {
				break
			}
		}
		for {
			j--
			if s[j] <= key {
				break
			}
		}
		if i < j {
			s[i], s[j] = s[j], s[i]
		}
	}
	arrayQuickSort(s, l, i-1)
	arrayQuickSort(s, j+1, r)
}

type linkedQuick struct {
	LinkedList
}

func NewLinkedQuick() LinkedSorter {
	return &linkedQuick{LinkedList{Type: "quickSort"}}
}

func (l linkedQuick) Sort(head *ListNode)  {
	if head == nil {
		panic("head is nil...")
	}
	tail := head
	for tail.Next != nil {
		tail =  tail.Next
	}
	linkListQuickSort(l.head, tail)
}

func linkListQuickSort(head, tail *ListNode) {

	if !reflect.DeepEqual(head, tail) && !reflect.DeepEqual(head.Next, tail) {
		mid := partition(head, tail)
		linkListQuickSort(head, mid)
		linkListQuickSort(mid.Next, tail)

	}
}

func partition(start, end *ListNode) *ListNode {
	p := start
	q := p.Next
	key := start.Val
	for q != end {
		if q.Val < key {
			p = p.Next
			p.Val, q.Val = q.Val, p.Val
		}
		q = q.Next
	}
	p.Val, start.Val = start.Val, p.Val
	return p
}