package main

func countServersDFS(grid [][]int) int {
	rows, col := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < col; j++ {
			ans += dfs(i, j, 0, grid)
		}
	}
	return ans
}

func dfs(i, j, count int, grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= rows || j >= cols || grid[i][j] == 0 {
		return count
	}
	count++
	grid[i][j] = 0
	for row := range rows {
		if grid[row][j] == 1 {
			count = dfs(row, j, count, grid)
		}
	}
	for col := range cols {
		if grid[i][col] == 1 {
			count = dfs(i, col, count, grid)
		}
	}
	if count == 1 {
		return 0
	}
	return count
}

func countServers(grid [][]int) int {
	rows, col := len(grid), len(grid[0])
	rowC, colC := make([]int, col), make([]int, rows)
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				rowC[j]++
				colC[i]++
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 && (rowC[j] > 1 || colC[i] > 1) {
				ans++
			}
		}
	}

	return ans
}

func main() {
	println(countServers([][]int{{1, 0}, {0, 1}}))    // 0
	println(countServersDFS([][]int{{1, 0}, {0, 1}})) // 0
	println(countServers([][]int{{1, 0}, {1, 1}}))    // 3
	println(countServersDFS([][]int{{1, 0}, {1, 1}})) // 3
}
