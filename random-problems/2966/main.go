package main

import (
	"fmt"
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums); i = i + 3 {
		if nums[i+2]-nums[i] > k {
			return [][]int{}
		}
		ans = append(ans, []int{nums[i], nums[i+1], nums[i+2]})
	}
	return ans
}

func main() {
	fmt.Println(divideArray([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2))                                 // [[1 1 3] [3 4 5] [7 8 9]]
	fmt.Println(divideArray([]int{2, 4, 2, 2, 5, 2}, 2))                                          // []
	fmt.Println(divideArray([]int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11}, 14)) // [[2 2 2] [4 5 5] [5 5 7] [7 8 8] [9 9 10] [11 12 12]]
}
