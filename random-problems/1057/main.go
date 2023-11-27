package main

import "fmt"

func assignBikes(workers [][]int, bikes [][]int) []int {
	dis := make([][][2]int, 2001)
	res := make([]int, len(workers))
	workersVisited := make([]bool, len(workers))
	bikesVisited := make([]bool, len(bikes))

	for i := 0; i < len(workers); i++ {
		for j := 0; j < len(bikes); j++ {
			distance := abs(workers[i][0]-bikes[j][0]) + abs(workers[i][1]-bikes[j][1])
			list := [2]int{i, j}
			dis[distance] = append(dis[distance], list)
		}
	}

	for i := 0; i < len(dis); i++ {
		for j := 0; j < len(dis[i]); j++ {
			if !workersVisited[dis[i][j][0]] && !bikesVisited[dis[i][j][1]] {
				res[dis[i][j][0]] = dis[i][j][1]
				workersVisited[dis[i][j][0]] = true
				bikesVisited[dis[i][j][1]] = true
			}
		}
	}
	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	workers := [][]int{{0, 0}, {1, 1}, {2, 0}}
	bikes := [][]int{{1, 0}, {2, 2}, {2, 1}}
	fmt.Println(assignBikes(workers, bikes))
}
