package main

import (
	"fmt"

	"github.com/iamsumit/go-dsa/ds/graph/g"
)

func main() {
	fmt.Println("--------------------------------")
	fmt.Println("Implementing Graph Data Structure")
	fmt.Println("--------------------------------")

	gVar := g.NewGraph()
	v1 := gVar.AddVertex(1)
	v2 := gVar.AddVertex(2)
	v3 := gVar.AddVertex(3)

	gVar.AddEdge(v1, v2)
	gVar.AddEdge(v2, v3)

	fmt.Println("Vertices:")
	for _, v := range gVar.GetVertices() {
		fmt.Printf("Vertex %d\n", v.Key)
	}

	fmt.Println("--------------------------------")

	fmt.Println("Adjacent Vertices:")
	for _, v := range gVar.GetVertices() {
		fmt.Printf("Vertex %d -> ", v.Key)
		for _, a := range gVar.GetAdjacent(v) {
			fmt.Printf("%d ", a.Key)
		}
		fmt.Println()
	}

	fmt.Println("--------------------------------")
}
