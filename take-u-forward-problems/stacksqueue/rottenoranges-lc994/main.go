package main

import "fmt"

// https://takeuforward.org/data-structure/rotten-oranges-min-time-to-rot-all-oranges-bfs/
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	queue := [][]int{}

	var countFresh int

	for row, rowV := range grid {
		for col, colV := range rowV {
			if colV == 2 {
				queue = append(queue, []int{row, col})
			}

			if colV == 1 {
				countFresh++
			}
		}
	}

	var count int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			points := queue[0]
			queue = queue[1:]

			rowCounter := []int{0, 0, 1, -1}
			colCounter := []int{1, -1, 0, 0}

			for incr := 0; incr < len(rowCounter); incr++ {
				row, col := points[0]+rowCounter[incr], points[1]+colCounter[incr]
				if row < 0 || col < 0 || row >= m || col >= n || grid[row][col] != 1 {
					continue
				}

				grid[row][col] = 2
				countFresh--
				queue = append(queue, []int{row, col})
			}
		}

		if len(queue) != 0 {
			count++
		}
	}

	if countFresh == 0 {
		return count
	}

	return -1
}

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	fmt.Println(orangesRotting([][]int{{0, 2}}))
	fmt.Println(orangesRotting([][]int{{2, 0, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 0, 1, 0, 0, 0, 0, 0, 0, 1}, {1, 0, 1, 0, 1, 1, 1, 1, 0, 1}, {1, 0, 1, 0, 1, 0, 0, 1, 0, 1}, {1, 0, 1, 0, 1, 0, 0, 1, 0, 1}, {1, 0, 1, 0, 1, 1, 0, 1, 0, 1}, {1, 0, 1, 0, 0, 0, 0, 1, 0, 1}, {1, 0, 1, 1, 1, 1, 1, 1, 0, 1}, {1, 0, 0, 0, 0, 0, 0, 0, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}))
}
