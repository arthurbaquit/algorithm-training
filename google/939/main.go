package main

import "math"

type point struct {
	x, y int
}

func minAreaRect(input [][]int) int {
	points, graph := make([]point, len(input)), map[point]bool{}
	for i, v := range input {
		points[i] = point{v[0], v[1]}
		graph[points[i]] = true
	}

	ans := math.MaxInt
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			if p1.x == p2.x || p1.y == p2.y {
				continue
			}
			s := abs(p1.x-p2.x) * abs(p1.y-p2.y)
			if s > ans {
				continue
			}
			if _, ok := graph[point{p1.x, p2.y}]; !ok {
				continue
			}
			if _, ok := graph[point{p2.x, p1.y}]; !ok {
				continue
			}
			ans = s
		}
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	println(minAreaRect([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}}))
	println(minAreaRect([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3}}))
	println(minAreaRect([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3}, {4, 4}}))
}
