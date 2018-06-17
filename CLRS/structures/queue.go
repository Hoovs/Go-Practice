package structures


type Queue struct {
	front, end *node
}

func InitQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(val int) {
	if q.front == nil {
		q.front = &node{val: val}
		q.end = q.front
	} else {
		q.end.next = &node{val: val}
		q.end = q.end.next
	}
}

func (q *Queue) Dequeue() int {
	if q.front == nil {
		panic("ops")
	} else {
		tmp := q.front
		q.front = q.front.next
		return tmp.val
	}
}