package main

import (
	"testing"

	"github.com/iamsumit/go-dsa/ds/linkedlist/ll"
)

func BenchmarkLinkedListInsert(b *testing.B) {
	l := &ll.LinkedList{}
	for i := 0; i < b.N; i++ {
		l.Insert(i)
	}
}

func BenchmarkLinkedListDeleteIn100(b *testing.B) {
	l := &ll.LinkedList{}
	for i := 0; i < 100; i++ {
		l.Insert(i)
	}
	for i := 0; i < b.N; i++ {
		l.Delete(56)
	}
}

func BenchmarkLinkedListDeleteIn1000(b *testing.B) {
	l := &ll.LinkedList{}
	for i := 0; i < 1000; i++ {
		l.Insert(i)
	}
	for i := 0; i < b.N; i++ {
		l.Delete(567)
	}
}

func BenchmarkLinkedListSearchIn100(b *testing.B) {
	l := &ll.LinkedList{}
	for i := 0; i < 100; i++ {
		l.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		l.Search(54)
	}
}

func BenchmarkLinkedListSearchIn1000(b *testing.B) {
	l := &ll.LinkedList{}
	for i := 0; i < 1000; i++ {
		l.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		l.Search(547)
	}
}

func BenchmarkLinkedListReverseIn100(b *testing.B) {
	l := &ll.LinkedList{}
	for i := 0; i < 100; i++ {
		l.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		l.Reverse()
	}
}

func BenchmarkLinkedListReverseIn1000(b *testing.B) {
	l := &ll.LinkedList{}
	for i := 0; i < 1000; i++ {
		l.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		l.Reverse()
	}
}
