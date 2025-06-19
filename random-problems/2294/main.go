package main

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	count := 1
	l, r := 0, 0
	for r < len(nums) {
		if nums[r]-nums[l] <= k {
			r++
			continue
		}
		count++
		l = r
		r++
	}
	return count
}

func main() {
	println(partitionArray([]int{3, 6, 1, 2, 5}, 2)) // 2
	println(partitionArray([]int{1, 2, 3}, 1))       // 2
	println(partitionArray([]int{2, 2, 4, 5}, 0))    // 3
}
