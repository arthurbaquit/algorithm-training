package main

import "fmt"

func distinctNumbers(nums []int, k int) []int {
	numberCount := make(map[int]int)
	distinctCount := 0
	for i := 0; i < k; i++ {
		numberCount[nums[i]]++
		if numberCount[nums[i]] == 1 {
			distinctCount++
		}
	}
	ans := []int{distinctCount}
	for i := 1; i < len(nums)-k+1; i++ {
		numberCount[nums[i-1]]--
		if numberCount[nums[i-1]] == 0 {
			distinctCount--
		}
		numberCount[nums[i+k-1]]++
		if numberCount[nums[i+k-1]] == 1 {
			distinctCount++
		}
		ans = append(ans, distinctCount)
	}
	return ans
}

func main() {
	fmt.Println(distinctNumbers([]int{1, 2, 3, 2, 2, 1, 3}, 3)) // [3,2,2,2,3]
	fmt.Println(distinctNumbers([]int{1, 1, 1, 1, 2, 3, 4}, 4)) // [1,2,3,4]
}
