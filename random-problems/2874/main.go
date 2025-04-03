package main

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	var ans int64
	maxLeft, maxRight := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		maxLeft[i] = max(maxLeft[i-1], nums[i-1])
		maxRight[n-i-1] = max(maxRight[n-i], nums[n-i])
	}
	for j := 1; j < n-1; j++ {
		ans = max(ans, int64(maxLeft[j]-nums[j])*int64(maxRight[j]))
	}
	return ans
}

func maximumTripletValueOnePass(nums []int) int64 {
	var ans int64
	maxDiff, maxI := 0, 0
	for _, num := range nums {
		ans = max(ans, int64(maxDiff)*int64(num))
		maxDiff = max(maxDiff, maxI-num)
		maxI = max(maxI, num)
	}
	return ans
}

func main() {
	println(maximumTripletValue([]int{12, 6, 1, 2, 7}))         // 77
	println(maximumTripletValue([]int{1, 10, 3, 4, 19}))        // 133
	println(maximumTripletValue([]int{1, 2, 3}))                // 0
	println(maximumTripletValueOnePass([]int{12, 6, 1, 2, 7}))  // 77
	println(maximumTripletValueOnePass([]int{1, 10, 3, 4, 19})) // 133
	println(maximumTripletValueOnePass([]int{1, 2, 3}))         // 0
}
