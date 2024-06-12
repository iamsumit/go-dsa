package main

import (
	"testing"

	"github.com/iamsumit/go-dsa/ds/heap/h"
)

func BenchmarkHeap_Insert(b *testing.B) {
	hs := h.Heap{}
	for i := 0; i < b.N; i++ {
		hs.Insert(i)
	}
}

func BenchmarkHeap_Delete(b *testing.B) {
	hs := h.Heap{}
	for i := 0; i < 10000; i++ {
		hs.Insert(i)
	}
	for i := 0; i < b.N; i++ {
		hs.Delete(i)
	}
}

func BenchmarkHeap_Peek(b *testing.B) {
	hs := h.Heap{}
	for i := 0; i < 10000; i++ {
		hs.Insert(i)
	}
	for i := 0; i < b.N; i++ {
		hs.Peek()
	}
}

func BenchmarkHeap_FindIndex(b *testing.B) {
	hs := h.Heap{}
	for i := 0; i < 10000; i++ {
		hs.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		hs.FindIndex(i)
	}
}

func BenchmarkHeap_UpdateVal(b *testing.B) {
	hs := h.Heap{}
	for i := 0; i < 10000; i++ {
		hs.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		rand := i % 10
		hs.UpdateVal(i, i+rand)
	}
}

func BenchmarkHeap_IncreaseByValue(b *testing.B) {
	hs := h.Heap{}
	for i := 0; i < 10000; i++ {
		hs.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		rand := i % 10
		hs.IncreaseByValue(i, rand)
	}
}

func BenchmarkHeap_DecreaseByValue(b *testing.B) {
	hs := h.Heap{}
	for i := 0; i < 10000; i++ {
		hs.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		rand := i % 10
		hs.DecreaseByValue(i, rand)
	}
}
