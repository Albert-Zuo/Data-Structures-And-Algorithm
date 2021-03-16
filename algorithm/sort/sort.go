package sort


import (
	"bytes"
	"fmt"
	"strconv"
)

type Sorter interface {
	Sort(s []int)
}

type Array struct {
	Type string
	slice []int
}

type ListNode struct {
	Val            int
	Next           *ListNode
}

type LinkedList struct {
	Type string
	head *ListNode
}

func (a Array) String() string {
	return fmt.Sprintf("%s: %v", a.Type ,a.slice)
}


func (l LinkedList) String() string {
	var (
		b bytes.Buffer
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
