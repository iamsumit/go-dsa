package main

import (
	"fmt"
)

// https://takeuforward.org/data-structure/m-coloring-problem/

func graphColoring(graph [][]int, m, n int) int {
	color := make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = 0
	}

	return solve(0, color, m, n, graph)
}

func solve(node int, color []int, m, n int, graph [][]int) int {
	if node == n {
		return 1
	}

	for i := 1; i <= m; i++ {
		if isSafe(node, color, graph, n, i) {
			color[node] = i
			if solve(node+1, color, m, n, graph) == 1 {
				return 1
			}

			color[node] = 0
		}
	}

	return 0
}

func isSafe(node int, color []int, graph [][]int, n, col int) bool {
	for i := 0; i < n; i++ {
		if i != node && graph[i][node] == 1 && color[i] == col {
			return false
		}
	}

	return true
}

func main() {
	// graph := make([][]int, 5)
	// for i := 0; i < 5; i++ {
	// 	graph[i] = make([]int, 5)
	// }
	// graph[0][1] = 1
	// graph[1][0] = 1
	// graph[1][2] = 1
	// graph[2][1] = 1
	// graph[2][3] = 1
	// graph[3][2] = 1
	// graph[3][0] = 1
	// graph[0][3] = 1
	// graph[0][2] = 1
	// graph[2][0] = 1
	// fmt.Println(graphColoring(graph, 3, 4))

	graph := make([][]int, 6)
	for i := 0; i < 6; i++ {
		graph[i] = make([]int, 6)
	}

	graph[5][1] = 1
	graph[5][3] = 1
	graph[4][5] = 1
	fmt.Println(graphColoring(graph, 1, 6))
}
