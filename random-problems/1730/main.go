package main

import "container/heap"

type Point struct {
	x, y, steps int
}

type MinHeap []Point

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i].steps < m[j].steps }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(Point))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	toPop := old[n-1]
	*m = old[0 : n-1]
	return toPop
}

func getFood(grid [][]byte) int {
	q := &MinHeap{}
	n, m := len(grid), len(grid[0])
	found := false
	for i := range n {
		if found {
			break
		}
		for j := range m {
			if grid[i][j] == '*' {
				heap.Push(q, Point{i, j, 0})
				grid[i][j] = 'X'
				found = true
				break
			}
		}
	}

	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	for q.Len() > 0 {
		curr := heap.Pop(q).(Point)
		for i := 0; i < 4; i++ {
			nextX, nextY := curr.x+dx[i], curr.y+dy[i]
			if nextX < 0 || nextX >= n || nextY < 0 || nextY >= m {
				continue
			}
			if grid[nextX][nextY] == 'X' {
				continue
			}
			if grid[nextX][nextY] == '#' {
				return curr.steps + 1
			}
			grid[nextX][nextY] = 'X'
			heap.Push(q, Point{nextX, nextY, curr.steps + 1})
		}
	}
	return -1
}

func main() {
	println(getFood([][]byte{{'X', 'X', 'X', 'X', 'X', 'X'}, {'X', '*', 'O', 'O', 'O', 'X'}, {'X', 'O', 'O', '#', 'O', 'X'}, {'X', 'X', 'X', 'X', 'X', 'X'}}))                                                                                   // 3
	println(getFood([][]byte{{'X', 'X', 'X', 'X', 'X'}, {'X', '*', 'X', 'O', 'X'}, {'X', 'O', 'X', '#', 'X'}, {'X', 'X', 'X', 'X', 'X'}}))                                                                                                       // -1
	println(getFood([][]byte{{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'}, {'X', '*', 'O', 'X', 'O', '#', 'O', 'X'}, {'X', 'O', 'O', 'X', 'O', 'O', 'X', 'X'}, {'X', 'O', 'O', 'O', 'O', '#', 'O', 'X'}, {'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'}})) // 6
}
