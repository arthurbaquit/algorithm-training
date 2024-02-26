package main

func maximumDetonation(bombs [][]int) int {
	neighbours := make(map[int][]int)
	for i := 0; i < len(bombs); i++ {
		x1 := bombs[i][0]
		y1 := bombs[i][1]
		r := bombs[i][2]
		for j := 0; j < len(bombs); j++ {
			if i != j {
				x2 := bombs[j][0]
				y2 := bombs[j][1]
				if (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) <= r*r {
					neighbours[i] = append(neighbours[i], j)
				}
			}
		}
	}
	max := 0
	for i := 0; i < len(bombs); i++ {
		visited := make(map[int]bool)
		q := []int{i}
		visited[i] = true
		for len(q) > 0 {
			first := q[0]
			q = q[1:]
			if val, ok := neighbours[first]; ok {
				for _, b := range val {
					if !visited[b] {
						visited[b] = true
						q = append(q, b)
					}
				}
			}
		}
		l := len(visited)
		if l == len(bombs) {
			return l
		}
		if l > max {
			max = l
		}
	}
	return max
}

func main() {
	println(maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}}))
	println(maximumDetonation([][]int{{1, 1, 5}, {10, 10, 5}}))
	println(maximumDetonation([][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}}))
}
