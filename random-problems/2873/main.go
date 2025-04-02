package main

func maximumTripletValue(nums []int) int64 {
	var ans int64
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]-nums[j] < 0 {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				ans = max(ans, int64(nums[i]-nums[j])*int64(nums[k]))
			}
		}
	}
	return ans
}

func maximumTripletValueGreedy(nums []int) int64 {
	var ans int64
	for k := 2; k < len(nums); k++ {
		maxI := nums[0]
		for j := 1; j < k; j++ {
			ans = max(ans, int64(maxI-nums[j])*int64(nums[k]))
			maxI = max(maxI, nums[j])
		}
	}

	return ans
}

func maximumTripletValuePrefixGreedy(nums []int) int64 {
	var ans int64
	n := len(nums)
	leftMax, rightMax := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], nums[i-1])
		rightMax[n-i-1] = max(rightMax[n-i], nums[n-i])
	}
	for k := 1; k < n-1; k++ {
		ans = max(ans, int64(leftMax[k]-nums[k])*int64(rightMax[k]))
	}

	return ans
}

func main() {
	println(maximumTripletValue([]int{12, 6, 1, 2, 7}))              // 77
	println(maximumTripletValue([]int{1, 10, 3, 4, 19}))             // 133
	println(maximumTripletValue([]int{1, 2, 3}))                     // 0
	println(maximumTripletValueGreedy([]int{12, 6, 1, 2, 7}))        // 77
	println(maximumTripletValueGreedy([]int{1, 10, 3, 4, 19}))       // 133
	println(maximumTripletValueGreedy([]int{1, 2, 3}))               // 0
	println(maximumTripletValuePrefixGreedy([]int{12, 6, 1, 2, 7}))  // 77
	println(maximumTripletValuePrefixGreedy([]int{1, 10, 3, 4, 19})) // 133
	println(maximumTripletValuePrefixGreedy([]int{1, 2, 3}))         // 0
}
