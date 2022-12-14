package main

import "fmt"

func sortedSquares(nums []int) []int {
	l := len(nums)
	squaredNumbers := make([]int, l)
	left, right := 0, l-1

	for i := l - 1; i >= 0; i-- {
		if nums[left]*nums[left] <= nums[right]*nums[right] {
			squaredNumbers[i] = nums[right] * nums[right]
			right--
			continue
		}
		squaredNumbers[i] = nums[left] * nums[left]
		left++
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
