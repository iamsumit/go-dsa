package main

import (
	"testing"

	"github.com/iamsumit/go-dsa/ds/graph/g"
)

func BenchmarkGraph(b *testing.B) {
	gVar := g.NewGraph()

	for i := 0; i < b.N; i++ {
		v1 := gVar.AddVertex(i)
		v2 := gVar.AddVertex(i)

		gVar.AddEdge(v1, v2)
	}
}

func BenchmarkGraphVertices100(b *testing.B) {
	gVar := g.NewGraph()

	for i := 0; i < 100; i++ {
		v1 := gVar.AddVertex(i)
		v2 := gVar.AddVertex(i)

		gVar.AddEdge(v1, v2)
	}

	for i := 0; i < b.N; i++ {
		for _, v := range gVar.GetVertices() {
			_ = v
		}
	}
}

func BenchmarkGraphVertices1000(b *testing.B) {
	gVar := g.NewGraph()

	for i := 0; i < 1000; i++ {
		v1 := gVar.AddVertex(i)
		v2 := gVar.AddVertex(i)

		gVar.AddEdge(v1, v2)
	}

	for i := 0; i < b.N; i++ {
		for _, v := range gVar.GetVertices() {
			_ = v
		}
	}
}

func BenchmarkGraphVertices10000(b *testing.B) {
	gVar := g.NewGraph()

	for i := 0; i < 10000; i++ {
		v1 := gVar.AddVertex(i)
		v2 := gVar.AddVertex(i)

		gVar.AddEdge(v1, v2)
	}

	for i := 0; i < b.N; i++ {
		for _, v := range gVar.GetVertices() {
			_ = v
		}
	}
}

func BenchmarkGraphAdjacent100(b *testing.B) {
	gVar := g.NewGraph()

	for i := 0; i < 100; i++ {
		v1 := gVar.AddVertex(i)
		v2 := gVar.AddVertex(i)

		gVar.AddEdge(v1, v2)
	}

	for i := 0; i < b.N; i++ {
		for _, v := range gVar.GetVertices() {
			for _, a := range gVar.GetAdjacent(v) {
				_ = a
			}
		}
	}
}

func BenchmarkGraphAdjacent1000(b *testing.B) {
	gVar := g.NewGraph()

	for i := 0; i < 1000; i++ {
		v1 := gVar.AddVertex(i)
		v2 := gVar.AddVertex(i)

		gVar.AddEdge(v1, v2)
	}

	for i := 0; i < b.N; i++ {
		for _, v := range gVar.GetVertices() {
			for _, a := range gVar.GetAdjacent(v) {
				_ = a
			}
		}
	}
}

func BenchmarkGraphAdjacent10000(b *testing.B) {
	gVar := g.NewGraph()

	for i := 0; i < 10000; i++ {
		v1 := gVar.AddVertex(i)
		v2 := gVar.AddVertex(i)

		gVar.AddEdge(v1, v2)
	}

	for i := 0; i < b.N; i++ {
		for _, v := range gVar.GetVertices() {
			for _, a := range gVar.GetAdjacent(v) {
				_ = a
			}
		}
	}
}
