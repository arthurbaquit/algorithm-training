package main

func findMaxFish(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			ans = max(ans, dfs(grid, i, j))
		}
	}
	return ans
}

func dfs(grid [][]int, x, y int) int {
	if grid[x][y] == 0 {
		return 0
	}
	maxDfs := grid[x][y]
	grid[x][y] = 0
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	for i := 0; i < 4; i++ {
		nextX, nextY := x+dx[i], y+dy[i]
		if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
			continue
		}
		if grid[nextX][nextY] == 0 {
			continue
		}
		maxDfs += dfs(grid, nextX, nextY)
	}
	return maxDfs
}

func main() {
	println(findMaxFish([][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}})) // 7
	println(findMaxFish([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}})) // 1
}
