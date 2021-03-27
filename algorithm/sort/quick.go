package sort

import (
	"reflect"
)

func init() {
	RegisterSorter("quickSort", &quick{})
}

// 快排
type quick struct {}



func(q *quick) Sort(s []int) {

	quickSort(s, 0, len(s)-1)
}

func quickSort(s []int, l, r int) {
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
	quickSort(s, l, i-1)
	quickSort(s, j+1, r)
}

type linkedQuick struct {
	LinkedList
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