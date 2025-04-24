package main

func countCompleteSubarraysBF(nums []int) int {
	n := len(nums)
	ans := 0
	distinctMap := make(map[int]bool)
	for i := range n {
		distinctMap[nums[i]] = true
	}
	distinctCount := len(distinctMap)
	mainDisCount := distinctCount
	for i := range n {
		distinctCount = 0
		distinctMap = make(map[int]bool)
		for j := i; j < n; j++ {
			if _, ok := distinctMap[nums[j]]; !ok {
				distinctCount++
				distinctMap[nums[j]] = true
			}
			if distinctCount == mainDisCount {
				ans++
			}
		}
	}
	return ans
}

func countCompleteSubarrays(nums []int) int {
	freq := make(map[int]int)
	n := len(nums)
	ans := 0
	distinctMap := make(map[int]bool)
	for _, num := range nums {
		distinctMap[num] = true
	}
	distinctCount := len(distinctMap)
	right := 0
	for i := range n {
		if i > 0 {
			freq[nums[i-1]]--
			if freq[nums[i-1]] == 0 {
				delete(freq, nums[i-1])
			}
		}
		for right < n && len(freq) < distinctCount {
			freq[nums[right]]++
			right++
		}
		if len(freq) == distinctCount {
			ans += (n - right + 1)
		}
	}
	return ans
}

func main() {
	println(countCompleteSubarraysBF([]int{1, 3, 1, 2, 2})) // 4
	println(countCompleteSubarraysBF([]int{5, 5, 5, 5}))    // 10
	println(countCompleteSubarrays([]int{1, 3, 1, 2, 2}))   // 4
	println(countCompleteSubarrays([]int{5, 5, 5, 5}))      // 10
}
