package main

import "fmt"

func queryResults(_ int, queries [][]int) []int {
	colorMap := make(map[int]int)
	ballsColor := make(map[int]int)
	distinctColors := 0
	ans := make([]int, len(queries))
	for i, query := range queries {
		ballToColor, color := query[0], query[1]
		if ballsColor[ballToColor] != 0 {
			prevColor := ballsColor[ballToColor]
			colorMap[prevColor]--
			if colorMap[prevColor] == 0 {
				distinctColors--
			}
		}
		ballsColor[ballToColor] = color
		colorMap[color]++
		if colorMap[color] == 1 {
			distinctColors++
		}
		ans[i] = distinctColors
	}
	return ans
}

func main() {
	fmt.Println(queryResults(3, [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}}))         // [1, 2, 2, 3]
	fmt.Println(queryResults(3, [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}})) // [1, 2, 2, 3, 4]
}
