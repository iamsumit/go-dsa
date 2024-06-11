package q

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}

	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[0]
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) Clear() {
	q.data = []int{}
}

func (q *Queue) Values() []int {
	return q.data
}
