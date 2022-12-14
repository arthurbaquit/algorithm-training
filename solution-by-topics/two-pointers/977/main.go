package main

import "fmt"

func sortedSquares(nums []int) []int {
	l := len(nums)
	squaredNumbers := make([]int, l)
	index := 0
	left, right := 0, 0
	for right < l && nums[right] < 0 {
		right++
	}
	left = right - 1
	for left >= 0 && right < l {
		if nums[left]*nums[left] <= nums[right]*nums[right] {
			squaredNumbers[index] = nums[left] * nums[left]
			index++
			left--
			continue
		}
		squaredNumbers[index] = nums[right] * nums[right]
		index++
		right++
	}
	for left >= 0 {
		squaredNumbers[index] = nums[left] * nums[left]
		index++
		left--
	}
	for right < l {
		squaredNumbers[index] = nums[right] * nums[right]
		index++
		right++
	}
	return squaredNumbers
}

func main() {
	sorted := sortedSquares([]int{-4, -1, 0, 3, 10})
	for _, v := range sorted {
		fmt.Printf("%v ", v)
	}
	println()
	sorted = sortedSquares([]int{-7, -3, 2, 3, 11})
	for _, v := range sorted {
		fmt.Printf("%v ", v)
	}
	println()

	sorted = sortedSquares([]int{-100, -99, -80})
	for _, v := range sorted {
		fmt.Printf("%v ", v)
	}
}
