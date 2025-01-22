package main

import (
	"container/heap"
	"fmt"
)

type Point struct {
	x, y, height int
}

type MinHeap []Point

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i].height < m[j].height }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	toPop := old[n-1]
	*m = old[0 : n-1]
	return toPop
}
func (m *MinHeap) Push(x any) { *m = append(*m, x.(Point)) }

func highestPeak(isWater [][]int) [][]int {
	ans := make([][]int, len(isWater))
	n, m := len(isWater), len(isWater[0])
	for i := range n {
		ans[i] = make([]int, m)
	}
	q := &MinHeap{}
	dirI := []int{0, 0, 1, -1}
	dirJ := []int{1, -1, 0, 0}
	for i := range n {
		for j := range m {
			ans[i][j] = -1
			if isWater[i][j] == 1 {
				ans[i][j]++
				heap.Push(q, Point{i, j, ans[i][j]})
			}
		}
	}
	for q.Len() > 0 {
		curr := heap.Pop(q).(Point)
		for k := range 4 {
			nextI, nextJ := curr.x+dirI[k], curr.y+dirJ[k]
			if !isValid(nextI, nextJ, n, m) {
				continue
			}
			if isWater[nextI][nextJ] == 1 || ans[nextI][nextJ] != -1 {
				continue
			}
			ans[nextI][nextJ] = curr.height + 1
			heap.Push(q, Point{nextI, nextJ, ans[nextI][nextJ]})
		}
	}
	return ans
}

func isValid(i, j, n, m int) bool {
	return i >= 0 && j >= 0 && i < n && j < m
}

func main() {
	fmt.Println(highestPeak([][]int{{0, 1}, {0, 0}}))                  // [[1,0],[2,1]]
	fmt.Println(highestPeak([][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}})) // [[1,1,0],[0,1,1],[1,2,2]]
}
