package main

func firstCompleteIndex(arr []int, mat [][]int) int {
	n, m := len(mat), len(mat[0])

	countR := make([]int, n)
	countC := make([]int, m)
	mapMat := make(map[int][]int)
	for i := range n {
		for j := range m {
			mapMat[mat[i][j]] = []int{i, j}
		}
	}
	for idx, v := range arr {
		i, j := mapMat[v][0], mapMat[v][1]
		countR[i]++
		countC[j]++
		if countR[i] == m || countC[j] == n {
			return idx
		}
	}
	return 0
}

func main() {
	println(firstCompleteIndex([]int{1, 3, 4, 2}, [][]int{{1, 4}, {2, 3}}))                                 // 2
	println(firstCompleteIndex([]int{2, 8, 7, 4, 1, 3, 5, 6, 9}, [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}})) // 3
}
