package modal

type Queue struct {
	q		[]string
	length	int
}

func NewQueue() *Queue {
	queue := &Queue{
		q: make([]string, 0, 2000),
	}
	return queue
}

func (q *Queue) Push(i string) {
	q.q = append(q.q, i)
	q.length += 1
}

func (q *Queue) Shift() (string) {
	item := q.q[0]
	q.q = q.q[1:]
	return item
}

func (q *Queue) Len() (int) {
	return q.length
}

func (q *Queue) HasValue(v string) (bool) {
	for _, value := range q.q {
		if (value == v) {
			return true
		}
	}
	return false
}

