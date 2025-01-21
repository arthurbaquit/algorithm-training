package main

func gridGame(grid [][]int) int64 {
	m := len(grid[0])
	upRowS := 0
	for j := range m {
		upRowS += grid[0][j]
	}
	downRowS := 0
	var minSum int64
	// this is the max sum for the problem
	// so it can be used as int64.Inf for this case
	minSum = 100000 * 5 * 10000
	for wentDownIdx := range m {
		upRowS -= grid[0][wentDownIdx]
		minSum = min(minSum, int64(max(upRowS, downRowS)))
		downRowS += grid[1][wentDownIdx]
	}
	return minSum
}

func main() {
	println(gridGame([][]int{{2, 5, 4}, {1, 5, 1}}))        // 4
	println(gridGame([][]int{{3, 3, 1}, {8, 5, 2}}))        // 4
	println(gridGame([][]int{{1, 3, 1, 15}, {1, 3, 3, 1}})) // 7
}
