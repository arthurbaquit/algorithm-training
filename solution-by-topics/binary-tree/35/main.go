package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid
			continue
		}
		left = mid + 1
	}
	return left
}

func main() {
	println(searchInsert([]int{1, 3, 5, 6}, 5))
	println(searchInsert([]int{1, 3, 5, 6}, 2))
	println(searchInsert([]int{1, 3, 5, 6}, 7))
	println(searchInsert([]int{1, 3, 5, 6}, 0))
	println(searchInsert([]int{1}, 0))
}
