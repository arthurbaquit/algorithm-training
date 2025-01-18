package main

import "container/heap"

type Point struct {
	x, y, cost int
}

type MinHeap []Point

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Point))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// Distance array to store minimum cost to reach each cell
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = 1e9 // Initialize value higher than nxm
		}
	}
	// Starting point has 0 cost
	dist[0][0] = 0

	// Directions: Right (1), Left (2), Down (3), Up (4)
	directions := []struct{ dx, dy int }{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	// Min heap to process the cell with the smallest cost first
	h := &MinHeap{}
	heap.Push(h, Point{0, 0, 0}) // Push starting point (0, 0) with 0 cost

	// Dijkstra's algorithm
	for h.Len() > 0 {
		// Pop the cell with the smallest cost
		cell := heap.Pop(h).(Point)
		x, y, cost := cell.x, cell.y, cell.cost

		// If we have already visited this cell, continue
		if cost > dist[x][y] {
			continue
		}

		// Check all possible directions
		for i := 0; i < 4; i++ {
			// Calculate the next cell's coordinates
			newX, newY := x+directions[i].dx, y+directions[i].dy
			if newX < 0 || newY < 0 || newX >= m || newY >= n {
				continue
			}

			// Calculate the new cost to reach this cell
			newCost := cost

			// if the direction does not match, update cost
			if grid[x][y] != i+1 {
				newCost++
			}

			// If the new cost is better, update it and push to the heap
			if newCost < dist[newX][newY] {
				dist[newX][newY] = newCost
				heap.Push(h, Point{newX, newY, newCost})
			}
		}
	}

	// Return the minimum cost to reach the bottom-right cell
	return dist[m-1][n-1]
}

func main() {
	grid := [][]int{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{1, 1, 1, 1},
		{2, 2, 2, 2},
	}
	println(minCost(grid)) // 3
	grid = [][]int{
		{1, 1, 3},
		{3, 2, 2},
		{1, 1, 4},
	}
	println(minCost(grid)) // 0
	grid = [][]int{
		{1, 2},
		{4, 3},
	}
	println(minCost(grid)) // 1

}
