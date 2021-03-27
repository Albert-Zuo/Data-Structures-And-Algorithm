package sort

import (
	"bytes"
	"strconv"
)

type Sorter interface {
	Sort([]int)
}


type ListNode struct {
	Val        int
	Next, Prev *ListNode
}

type LinkedList struct {
	Type string
	head *ListNode
}

func (l LinkedList) String() string {
	var (
		b    bytes.Buffer
		curr = l.head
	)
	b.WriteString(l.Type + ": [")
	for curr != nil {
		b.WriteString(strconv.Itoa(curr.Val))
		if curr.Next != nil {
			b.WriteString(" -> ")
		}
		curr = curr.Next
	}
	b.WriteString("]")
	return b.String()
}

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}