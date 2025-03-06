package main

import "fmt"

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	arr := make([]bool, n*n+1)
	ans := make([]int, 2)
	for _, r := range grid {
		for _, v := range r {
			if arr[v] {
				ans[0] = v
			}
			arr[v] = true
		}
	}
	for i := 1; i <= n*n; i++ {
		if !arr[i] {
			ans[1] = i
			return ans
		}
	}
	return ans
}

func main() {
	fmt.Println(findMissingAndRepeatedValues([][]int{{1, 3}, {2, 2}}))                  // [2, 4]
	fmt.Println(findMissingAndRepeatedValues([][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}})) // [9,5]
}
