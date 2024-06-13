package main

import (
	"testing"

	"github.com/iamsumit/go-dsa/ds/queue/pq"
	"github.com/iamsumit/go-dsa/ds/queue/q"
)

func BenchmarkEnqueue(b *testing.B) {
	_q := q.Queue{}
	for i := 0; i < b.N; i++ {
		_q.Enqueue(i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	_q := q.Queue{}

	for i := 0; i < 10000; i++ {
		_q.Enqueue(i)
	}

	for i := 0; i < b.N; i++ {
		_q.Dequeue()
	}
}

func BenchmarkClear(b *testing.B) {
	_q := q.Queue{}

	for i := 0; i < 10000; i++ {
		_q.Enqueue(i)
	}

	for i := 0; i < b.N; i++ {
		_q.Clear()
	}
}

func BenchmarkPeek(b *testing.B) {
	_q := q.Queue{}

	for i := 0; i < 10000; i++ {
		_q.Enqueue(i)
	}

	for i := 0; i < b.N; i++ {
		_q.Peek()
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	_q := q.Queue{}

	for i := 0; i < b.N; i++ {
		_q.IsEmpty()
	}
}

func BenchmarkSize(b *testing.B) {
	_q := q.Queue{}

	for i := 0; i < 10000; i++ {
		_q.Enqueue(i)
	}

	for i := 0; i < b.N; i++ {
		_q.Size()
	}
}

func BenchmarkPriorityQueueEnqueue(b *testing.B) {
	_pq := pq.PriorityQueue{}

	for i := 0; i < b.N; i++ {
		_pq.Enqueue(i, i)
	}
}

func BenchmarkPriorityQueueWith09Priority(b *testing.B) {
	_pq := pq.PriorityQueue{}

	for i := 0; i < b.N; i++ {
		priority := i % 10
		_pq.Enqueue(i, priority)
	}
}

func BenchmarkPriorityQueueDequeue(b *testing.B) {
	_pq := pq.PriorityQueue{}

	for i := 0; i < 10000; i++ {
		_pq.Enqueue(i, i)
	}

	for i := 0; i < b.N; i++ {
		_pq.Dequeue()
	}
}

func BenchmarkPriorityQueueClear(b *testing.B) {
	_pq := pq.PriorityQueue{}

	for i := 0; i < 10000; i++ {
		_pq.Enqueue(i, i)
	}

	for i := 0; i < b.N; i++ {
		_pq.Clear()
	}
}

func BenchmarkPriorityQueuePeek(b *testing.B) {
	_pq := pq.PriorityQueue{}

	for i := 0; i < 10000; i++ {
		_pq.Enqueue(i, 0)
	}

	for i := 0; i < b.N; i++ {
		_pq.Peek()
	}
}

func BenchmarkPriorityQueueIsEmpty(b *testing.B) {
	_pq := pq.PriorityQueue{}

	for i := 0; i < b.N; i++ {
		_pq.IsEmpty()
	}
}

func BenchmarkPriorityQueueSize(b *testing.B) {
	_pq := pq.PriorityQueue{}

	for i := 0; i < 10000; i++ {
		_pq.Enqueue(i, i)
	}

	for i := 0; i < b.N; i++ {
		_pq.Size()
	}
}
