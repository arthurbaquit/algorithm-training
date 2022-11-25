package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	l, i, j, k := len(nums), 2, 1, 0
	for i < l && j < l && k < l {
		if nums[i]+nums[j] <= nums[k] {
			k++
			j++
			i++
			continue
		}
		return nums[i] + nums[j] + nums[k]
	}
	return 0
}

func main() {
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}))
	fmt.Println(largestPerimeter([]int{3, 2, 3, 4}))
	fmt.Println(largestPerimeter([]int{1, 2, 1}))
}
