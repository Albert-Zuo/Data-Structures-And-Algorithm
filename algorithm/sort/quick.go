package sort

import (
	"reflect"
)

// 快排
type quick struct {
	Array
}



func NewQuickSort() Sorter {
	return &quick{Array{Type: "quickSort"}}
}

func(q *quick) Sort(s []int) {
	q.slice = s
	quickSort(q.slice[:], 0, len(q.slice)-1)
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





//head ,next := &sort.ListNode{Val: 44}, &sort.ListNode{Val: 6}
//head.Next = next
//
//prev, tail := &sort.ListNode{Val: 17}, &sort.ListNode{Val: 15}
//next.Next = &sort.ListNode{Val: 6, Next: prev}
//prev.Next = tail
//sort.QuickSortLinkList(head, tail)
//curr := head
//for i := 0; i < 5; i++ {
//	fmt.Println(curr.Val)
//	curr = curr.Next
//}
func QuickSortLinkList(head, tail *ListNode) {

	if !reflect.DeepEqual(head, tail) && !reflect.DeepEqual(head.Next, tail) {
		mid := partition(head, tail)
		QuickSortLinkList(head, mid)
		QuickSortLinkList(mid.Next, tail)

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