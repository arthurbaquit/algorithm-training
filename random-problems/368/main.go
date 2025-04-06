package main

import (
	"fmt"
	"slices"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([]int, n)
	hash := make([]int, n)
	for i := range dp {
		dp[i] = 1
		hash[i] = i
	}
	maxI, last := 1, 0

	for i := range n {
		for j := range i {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				hash[i] = j
			}
		}
		if dp[i] > maxI {
			maxI = dp[i]
			last = i
		}
	}

	res := []int{}
	for hash[last] != last {
		res = append(res, nums[last])
		last = hash[last]
	}
	res = append(res, nums[last])
	slices.Reverse(res)
	return res
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{1, 2, 3}))    // [1, 2])
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 8})) // [1, 2, 4, 8])
}
