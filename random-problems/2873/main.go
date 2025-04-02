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

func main() {
	println(maximumTripletValue([]int{12, 6, 1, 2, 7}))  // 77
	println(maximumTripletValue([]int{1, 10, 3, 4, 19})) // 133
	println(maximumTripletValue([]int{1, 2, 3}))         // 0
}
