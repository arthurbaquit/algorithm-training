package main

import (
	"fmt"
	"math"
)

var directions = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	var queue [][2]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			} else {
				mat[i][j] = math.MaxInt64
			}
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			x, y := curr[0]+dir[0], curr[1]+dir[1]
			if x < 0 || x >= m || y < 0 || y >= n || mat[x][y] <= mat[curr[0]][curr[1]] {
				continue
			}
			queue = append(queue, [2]int{x, y})
			mat[x][y] = mat[curr[0]][curr[1]] + 1
		}
	}

	return mat
}

func updateMatrixWithDP(mat [][]int) [][]int {

	rows := len(mat)
	if rows == 0 {
		return mat
	}
	cols := len(mat[0])

	//First check
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 1 {
				mat[i][j] = 10000
				if i > 0 {
					min(&mat[i][j], mat[i-1][j]+1)
				}
				if j > 0 {
					min(&mat[i][j], mat[i][j-1]+1)
				}
			}
		}
	}

	//Second check
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if i < rows-1 {
				min(&mat[i][j], mat[i+1][j]+1)
			}
			if j < cols-1 {
				min(&mat[i][j], mat[i][j+1]+1)
			}
		}
	}

	return mat
}

func min(d1 *int, d2 int) {
	if *d1 > d2 {
		*d1 = d2
	}
}
func main() {
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(updateMatrix([][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 0}}))
	fmt.Println(updateMatrixWithDP([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
	fmt.Println(updateMatrixWithDP([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(updateMatrixWithDP([][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 0}}))
	
}
