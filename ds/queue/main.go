package main

import (
	"fmt"

	"github.com/iamsumit/go-dsa/ds/queue/pq"
	"github.com/iamsumit/go-dsa/ds/queue/q"
)

func main() {
	fmt.Println("-------------------------")
	fmt.Println("Queue implementation in Go")
	fmt.Println("-------------------------")

	_q := q.Queue{}
	fmt.Println("Initiated, data", _q.Values())
	_q.Enqueue(1)
	fmt.Println("Enqueued 1, data", _q.Values())
	_q.Enqueue(2)
	fmt.Println("Enqueued 2, data", _q.Values())

	fmt.Println("-------------------------")

	fmt.Println("Peek:", _q.Peek())
	fmt.Println("Size:", _q.Size())
	fmt.Println("Dequeue:", _q.Dequeue())
	fmt.Println("Peek:", _q.Peek())
	fmt.Println("Data:", _q.Values())

	fmt.Println("-------------------------")

	fmt.Println("IsEmpty:", _q.IsEmpty())
	fmt.Println("Dequeue:", _q.Dequeue())
	fmt.Println("IsEmpty:", _q.IsEmpty())

	fmt.Println("-------------------------")

	_q.Enqueue(1)
	_q.Enqueue(2)

	fmt.Println("Enqueued 1, 2, data", _q.Values())
	_q.Clear()
	fmt.Println("Cleared, data", _q.Values())

	fmt.Println("-------------------------")

	fmt.Println("Priority Queue implementation in Go")

	fmt.Println("-------------------------")

	_pq := pq.PriorityQueue{}
	fmt.Println("Initiated, data", _pq.Values())
	_pq.Enqueue(1, 1)
	fmt.Println("Enqueued 1, data", _pq.Values())
	_pq.Enqueue(2, 2)
	fmt.Println("Enqueued 2, data", _pq.Values())
	_pq.Enqueue(3, 0)
	fmt.Println("Enqueued 2, data", _pq.Values())

	fmt.Println("-------------------------")

	fmt.Println("Peek:", _pq.Peek())
	fmt.Println("Size:", _pq.Size())
	fmt.Println("Dequeue:", _pq.Dequeue())
	fmt.Println("Peek:", _pq.Peek())
	fmt.Println("Data:", _pq.Values())

	fmt.Println("-------------------------")

	fmt.Println("IsEmpty:", _pq.IsEmpty())
	fmt.Println("Dequeue:", _pq.Dequeue())
	fmt.Println("Dequeue:", _pq.Dequeue())
	fmt.Println("IsEmpty:", _pq.IsEmpty())

	fmt.Println("-------------------------")

	_pq.Enqueue(1, 1)
	_pq.Enqueue(2, 2)

	fmt.Println("Enqueued 1, 2, data", _pq.Values())

	_pq.Clear()
	fmt.Println("Cleared, data", _pq.Values())

	fmt.Println("-------------------------")
}
