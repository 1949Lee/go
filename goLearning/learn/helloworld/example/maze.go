package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

func (p point) add(t point) point {
	return point{
		p.i + t.i,
		p.j + t.j,
	}
}
func (p point) locate(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

var dirs = []point{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range maze {
		steps[i] = make([]int, len(maze[0]))
	}
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.locate(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.locate(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			currStep, _ := cur.locate(steps)
			steps[next.i][next.j] = currStep + 1

			Q = append(Q, next)

		}
	}

	return steps
}

func getPath(steps [][]int, start, end point) []point {
	endStep, _ := end.locate(steps)
	P := []point{end}
	Q := []point{end}
	endStep--
	for endStep >= 0 {
		if endStep == 0 {
		}
		cur := Q[0]
		Q = Q[1:]
		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.locate(steps)
			if (!ok || val == 0) && next != start {
				continue
			}
			if val == endStep {
				P = append(P, next)
				Q = append(Q, next)
			}
		}
		endStep--
	}

	return P
}

func main() {
	maze := readMaze("example/maze.in")

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%4d ", val)
		}
		fmt.Println()
	}

	p := getPath(steps, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	fmt.Printf("共计%d步:\n", len(p))
	for i := len(p) - 1; i > 0; i-- {
		fmt.Printf("(%d, %d) ->", p[i].i+1, p[i].j+1)
	}

	fmt.Printf("(%d, %d)", p[0].i+1, p[0].j+1)
}
