package main

func countPairs(nums []int, k int) int {
	ans := 0
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i*j%k == 0 {
				ans++
			}
		}
	}
	return ans
}

func main() {
	println(countPairs([]int{3, 1, 2, 2, 2, 1, 3}, 2)) // 4
	println(countPairs([]int{1, 2, 3, 4}, 1))          // 0
}
