package main

func maxSubArray(nums []int) int {
	max := -100000
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > max {
			max = count
		}
		if count < 0 {
			count = 0
			continue
		}

	}
	return max
}

func main() {
	println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	println(maxSubArray([]int{1}))
	println(maxSubArray([]int{-1}))
	println(maxSubArray([]int{-2, -1}))
	println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
