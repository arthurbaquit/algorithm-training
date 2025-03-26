package main

import "sort"

func minOperations(grid [][]int, x int) int {
	n, m := len(grid), len(grid[0])
	numArr := make([]int, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			numArr[i*m+j] = grid[i][j]
		}
	}
	sort.Ints(numArr)
	result := 0
	median := numArr[n*m/2]
	for _, number := range numArr {
		if number%x != median%x {
			return -1
		}
		result += abs((median - number) / x)
	}
	return result
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	println(minOperations([][]int{{2, 4}, {6, 8}}, 2)) // 4
	println(minOperations([][]int{{1, 5}, {2, 3}}, 1)) // 5
	println(minOperations([][]int{{1, 2}, {3, 4}}, 2)) // -1
}
