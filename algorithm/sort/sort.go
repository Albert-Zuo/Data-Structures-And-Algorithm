package sort

import (
	"bytes"
	"fmt"
	"strconv"
)

type ArraySorter interface {
	Sort([]int)
}

type Array struct {
	Type  string
	slice []int
}

func (a Array) String() string {
	return fmt.Sprintf("%s: %v", a.Type, a.slice)
}

type LinkedSorter interface {
	Sort(head *ListNode)
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