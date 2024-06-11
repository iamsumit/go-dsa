package main

import (
	"testing"

	"github.com/iamsumit/go-dsa/ds/stack/st"
)

func BenchmarkStackPush(b *testing.B) {
	stack := st.Stack{}
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkStackPop(b *testing.B) {
	stack := st.Stack{}
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}

	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}

func BenchmarkStackPeek(b *testing.B) {
	stack := st.Stack{}
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}

	for i := 0; i < b.N; i++ {
		stack.Peek()
	}
}

func BenchmarkStackIsEmpty(b *testing.B) {
	stack := st.Stack{}

	for i := 0; i < b.N; i++ {
		stack.IsEmpty()
	}
}

func BenchmarkStackSize(b *testing.B) {
	stack := st.Stack{}
	for i := 0; i < b.N; i++ {
		stack.Size()
	}
}
