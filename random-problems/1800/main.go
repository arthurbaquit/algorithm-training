package main

func maxAscendingSum(nums []int) int {
	maxSum, currSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			currSum += nums[i]
			continue
		}
		maxSum = max(maxSum, currSum)
		currSum = nums[i]
	}
	return max(maxSum, currSum)
}

func main() {
	println(maxAscendingSum([]int{10, 20, 30, 5, 10, 50}))      // 65
	println(maxAscendingSum([]int{10, 20, 30, 40, 50}))         // 150
	println(maxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12})) // 33
}
