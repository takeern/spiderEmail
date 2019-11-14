package modal

type Queue struct {
	Q		[]string
	length	int
}

func NewQueue(len int) *Queue {
	queue := &Queue{
		Q: make([]string, 0, len),
	}
	return queue
}

func (q *Queue) Push(i string) {
	q.Q = append(q.Q, i)
	q.length += 1
}

func (q *Queue) PushList(i []string) {
	q.Q = append(q.Q, i...)
}

func (q *Queue) Shift() (string) {
	item := q.Q[0]
	q.Q = q.Q[1:]
	return item
}

func (q *Queue) Remove(v string) {
	for i, value := range q.Q {
		if (value == v) {
			q.Q = append(q.Q[:i], q.Q[i+1:]...)
			break;
		}
	}
}

func (q *Queue) Len() (int) {
	return len(q.Q)
}

func (q *Queue) HasValue(v string) (bool) {
	for _, value := range q.Q {
		if (value == v) {
			return true
		}
	}
	return false
}

