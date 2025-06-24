package main

import "fmt"

func findKDistantIndices(nums []int, key int, k int) []int {
	keys := []int{}
	for i, num := range nums {
		if num == key {
			keys = append(keys, i)
		}
	}
	ans := []int{}
	if len(keys) == 0 {
		return ans
	}
	for i := range nums {
		for _, ke := range keys {
			if abs(i-ke) <= k {
				ans = append(ans, i)
				break
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(findKDistantIndices([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1)) // [1, 2, 3, 4, 5, 6]
	fmt.Println(findKDistantIndices([]int{2, 2, 2, 2, 2}, 2, 2))       // [0, 1, 2, 3, 4]
}
