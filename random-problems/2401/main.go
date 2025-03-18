package main

func longestNiceSubarray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	maxLen := -1
	count := 0
	for i := range nums {
		count = 1
		for j := i + 1; j < len(nums); j++ {
			isNice := true
			for k := i; k < j; k++ {
				if nums[k]&nums[j] != 0 {
					isNice = false
					break
				}
			}
			if isNice {
				count++
				maxLen = max(maxLen, count)
				continue
			}
			break
		}
	}
	return max(maxLen, count)
}

func main() {
	println(longestNiceSubarray([]int{1, 3, 8, 48, 10})) // 3
	println(longestNiceSubarray([]int{3, 1, 5, 11, 13})) // 1
}
