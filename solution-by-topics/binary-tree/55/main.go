package main

func search(nums []int, target int) int {
	l := len(nums)
	left, right := 0, l-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
			continue
		}
		left = mid + 1
	}
	return -1
}

func main() {
	println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	println(search([]int{-1, 0, 3, 5, 9, 12}, 12))
	println(search([]int{-1, 0, 3, 5, 9, 12}, 13))
	println(search([]int{-1, 0, 3, 5, 9, 12}, -2))
	println(search([]int{-1, 0, 3, 5, 9, 12}, -1))
}
