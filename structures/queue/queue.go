package queue

import "errors"

var EmptyQueueErr = errors.New("queue is empty, can not dequeue")

type Queue interface {
	GetSize() int
	IsEmpty() bool
	Enqueue(e interface{})
	Dequeue() (interface{}, error)
	GetFront() (interface{}, error)
}
