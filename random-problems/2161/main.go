package main

import "fmt"

func pivotArray(nums []int, pivot int) []int {
	ans := make([]int, len(nums))
	l, r := 0, len(nums)-1
	i, j := 0, len(nums)-1
	for i < len(nums) {
		if nums[i] < pivot {
			ans[l] = nums[i]
			l++
		}
		if nums[j] > pivot {
			ans[r] = nums[j]
			r--
		}
		i++
		j--
	}
	for l <= r {
		ans[l] = pivot
		l++
	}
	return ans
}

func main() {
	fmt.Println(pivotArray([]int{9, 12, 5, 10, 14, 3, 10}, 10)) // [9 3 5 10 10 14 12]
	fmt.Println(pivotArray([]int{-3, 4, 3, 2}, 10))             //[-3 4 3 2]
}
