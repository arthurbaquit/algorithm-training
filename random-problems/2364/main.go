package main

func countBadPairs(nums []int) int64 {
	countDiff := make(map[int]int)
	var ans int64

	for i := range nums {
		countDiff[i-nums[i]]++
		ans += int64(i) - (int64(countDiff[i-nums[i]]) - int64(1))
	}

	return ans
}

func main() {
	println(countBadPairs([]int{1, 2, 3, 4, 5})) // 0
	println(countBadPairs([]int{1, 1, 1, 1, 1})) // 10
}
