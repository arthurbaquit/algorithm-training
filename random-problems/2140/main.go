package main

func mostPoints(questions [][]int) int64 {
	totalPointsFrom := make([]int64, len(questions))
	var maxPoints int64
	var currMax int64
	var getPoints func(int) int64
	getPoints = func(i int) int64 {
		if i >= len(questions) {
			return 0
		}
		v := totalPointsFrom[i]
		if v == 0 {
			tmp := int64(questions[i][0]) + getPoints(i+questions[i][1]+1)
			currMax = max(currMax, tmp)
			totalPointsFrom[i] = currMax
		}
		return totalPointsFrom[i]
	}
	for i := len(questions) - 1; i >= 0; i-- {
		maxPoints = max(maxPoints, getPoints(i))
	}
	return maxPoints
}

func main() {
	println(mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))         // 5
	println(mostPoints([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}})) // 7
}
