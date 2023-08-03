package main

import "fmt"

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	k = k % len(nums)
	idxToRotate := len(nums) - k
	aux := append(nums[idxToRotate:], nums[0:idxToRotate]...)
	for i, v := range aux {
		nums[i] = v
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
}
