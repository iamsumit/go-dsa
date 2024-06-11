package pq

type Data struct {
	Value    int
	Priority int
}

type PriorityQueue struct {
	data []Data
}

// We can either make Enqueue O(n) and Dequeue O(1) or Enqueue O(1) and Dequeue O(n)
func (pq *PriorityQueue) Enqueue(val, priority int) {
	data := Data{
		Value:    val,
		Priority: priority,
	}

	if pq.IsEmpty() {
		pq.data = append(pq.data, data)
		return
	}

	for i, d := range pq.data {
		if d.Priority > priority {
			pq.data = append(pq.data[:i], append([]Data{data}, pq.data[i:]...)...)
			return
		}
	}

	pq.data = append(pq.data, data)
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.data) == 0
}

func (pq *PriorityQueue) Dequeue() int {
	if pq.IsEmpty() {
		return -1
	}

	val := pq.data[0].Value

	pq.data = pq.data[1:]

	return val
}

func (pq *PriorityQueue) Peek() int {
	if pq.IsEmpty() {
		return -1
	}

	return pq.data[0].Value
}

func (pq *PriorityQueue) Size() int {
	return len(pq.data)
}

func (pq *PriorityQueue) Clear() {
	pq.data = []Data{}
}

func (pq *PriorityQueue) Values() []Data {
	return pq.data
}

func (pq *PriorityQueue) Priorities() []int {
	priorities := make([]int, len(pq.data))
	for i, d := range pq.data {
		priorities[i] = d.Priority
	}

	return priorities
}
