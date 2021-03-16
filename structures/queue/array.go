package queue

type ArrayQueue struct {
	data []interface{}
}

func NewArrayQueue() Queue {
	return &ArrayQueue{make([]interface{}, 0)}
}

func (q *ArrayQueue) GetSize() int {
	return len(q.data)
}
func (q *ArrayQueue) IsEmpty() bool {
	return len(q.data) == 0
}
func (q *ArrayQueue) Enqueue(e interface{}) {
	q.data = append(q.data, e)
}
func (q *ArrayQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, EmptyQueueErr
	}
	e := q.data[0]
	q.data = q.data[1:]
	return e, nil
}

func (q *ArrayQueue) GetFront() (interface{}, error) {
	if q.IsEmpty() {
		return nil, EmptyQueueErr
	}
	e := q.data[0]
	return e, nil
}
