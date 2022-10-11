package main

import "fmt"

const MaxInt32 = 1<<31 - 1

func increasingTriplet(nums []int) bool {
	l := len(nums)
	if l < 3 {
		return false
	}
	min, mid := MaxInt32, MaxInt32
	for k := 0; k < l; k++ {
		if nums[k] > mid {
			return true
		}
		if nums[k] > min && nums[k] < mid {
			mid = nums[k]
			continue
		}
		if nums[k] < min {
			min = nums[k]
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(increasingTriplet(nums))
	nums = []int{5, 4, 3, 2, 1}
	fmt.Println(increasingTriplet(nums))
	nums = []int{2, 1, 5, 0, 4, 6}
	fmt.Println(increasingTriplet(nums))
	nums = []int{1, 2, 3, 1, 2, 1}
	fmt.Println(increasingTriplet(nums))
}
