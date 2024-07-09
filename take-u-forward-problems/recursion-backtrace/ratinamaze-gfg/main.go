package main

import "fmt"

// https://takeuforward.org/data-structure/rat-in-a-maze/
func ratInAMaze(maze [][]int) []string {
	if maze[0][0] == 0 {
		return []string{}
	}

	directions := []string{}
	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze))
	}

	var directionFunc func(row, col int, dir string)
	directionFunc = func(row, col int, dir string) {
		if row == len(maze)-1 && col == len(maze)-1 {
			directions = append(directions, dir)
			return
		}

		visited[row][col] = true

		if col-1 >= 0 && maze[row][col-1] == 1 && !visited[row][col-1] {
			directionFunc(row, col-1, dir+"L")
		}

		if row+1 < len(maze) && maze[row+1][col] == 1 && !visited[row+1][col] {
			directionFunc(row+1, col, dir+"D")
		}

		if col+1 < len(maze) && maze[row][col+1] == 1 && !visited[row][col+1] {
			directionFunc(row, col+1, dir+"R")
		}

		if row-1 >= 0 && maze[row-1][col] == 1 && !visited[row-1][col] {
			directionFunc(row-1, col, dir+"U")
		}

		visited[row][col] = false
	}

	directionFunc(0, 0, "")

	return directions
}

func main() {
	// fmt.Println(ratInAMaze([][]int{
	// 	{1, 0, 0, 0},
	// 	{1, 1, 0, 1},
	// 	{1, 1, 0, 0},
	// 	{0, 1, 1, 1},
	// }))

	// fmt.Println(ratInAMaze([][]int{
	// 	{1, 1, 1, 1},
	// 	{1, 1, 1, 1},
	// 	{1, 1, 1, 1},
	// 	{1, 1, 1, 1},
	// }))

	// fmt.Println(ratInAMaze([][]int{
	// 	{1, 1, 1, 1, 1},
	// 	{1, 1, 1, 1, 1},
	// 	{1, 1, 1, 1, 1},
	// 	{1, 1, 1, 1, 1},
	// 	{1, 1, 1, 1, 1},
	// }))

	fmt.Println(ratInAMaze([][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 0},
		{1, 0, 1, 1},
		{0, 0, 1, 1},
	}))

	// DDDDRRRULLURULURRRDDDD
	// DDDDRRRULLURULURRRDLDRDD - incorrect

	// fmt.Println(ratInAMaze([][]int{
	// 	{1, 0},
	// 	{1, 0},
	// }))
}
