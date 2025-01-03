package main

func waysToSplitArray(nums []int) int {
	totalSum := 0
	l := len(nums)
	for _, n := range nums {
		totalSum += n
	}
	count := 0
	currSum := 0
	for i := 0; i < l-1; i++ {
		n := nums[i]
		currSum += n
		if currSum < totalSum-currSum {
			continue
		}
		count++
	}
	return count
}

func main() {
	println(waysToSplitArray([]int{10, 4, -8, 7}))
	println(waysToSplitArray([]int{2, 3, 1, 0}))
}
