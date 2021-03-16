package queue

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

type LinkedQueue struct {
	head, tail *ListNode
	size       int
}

func NewLinkedQueue() Queue {
	return &LinkedQueue{}
}

func (l *LinkedQueue) GetSize() int {
	return l.size
}

func (l *LinkedQueue) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedQueue) Enqueue(e interface{}) {
	node := &ListNode{Val: e}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		l.tail = node
	}
}

func (l *LinkedQueue) Dequeue() (interface{}, error) {
	if l.IsEmpty() {
		return nil, EmptyQueueErr
	}
	e := l.head
	l.head = l.head.Next
	return e.Val, nil
}

func (l *LinkedQueue) GetFront() (interface{}, error) {
	if l.IsEmpty() {
		return nil, EmptyQueueErr
	}
	return l.head.Val, nil
}
