package main

import "container/heap"

type MinHeap []Point
type Point struct {
	x, y, minH int
}

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i].minH < m[j].minH }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(Point))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}
func trapRainWater(heightMap [][]int) int {
	numRows, numCols := len(heightMap), len(heightMap[0])
	isVisited := make([][]bool, numRows)
	for i := range isVisited {
		isVisited[i] = make([]bool, numCols)
	}
	totalWater := 0
	boundary := &MinHeap{}
	for j := 0; j < numCols; j++ {
		heap.Push(boundary, Point{0, j, heightMap[0][j]})
		heap.Push(boundary, Point{numRows - 1, j, heightMap[numRows-1][j]})
		isVisited[0][j] = true
		isVisited[numRows-1][j] = true
	}
	for i := 1; i < numRows-1; i++ {
		heap.Push(boundary, Point{i, 0, heightMap[i][0]})
		heap.Push(boundary, Point{i, numCols - 1, heightMap[i][numCols-1]})
		isVisited[i][0] = true
		isVisited[i][numCols-1] = true
	}
	directions := []struct{ dx, dy int }{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for boundary.Len() > 0 {
		curr := heap.Pop(boundary).(Point)
		currR := curr.x
		currC := curr.y
		currH := curr.minH
		for _, dir := range directions {
			neighborR := currR + dir.dx
			neighborC := currC + dir.dy
			if neighborR < 0 || neighborR >= numRows || neighborC < 0 || neighborC >= numCols || isVisited[neighborR][neighborC] {
				continue
			}
			neighborH := heightMap[neighborR][neighborC]
			if neighborH < currH {
				totalWater += currH - neighborH
			}
			heap.Push(boundary, Point{neighborR, neighborC, max(currH, neighborH)})
			isVisited[neighborR][neighborC] = true
		}
	}
	return totalWater

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(trapRainWater([][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}))                          // 4
	println(trapRainWater([][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}})) // 10
}
