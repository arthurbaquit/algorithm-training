package main

func MyLongestMonotonicSubarray(nums []int) int {
	maxStreak := 1
	currStreak := 1
	isInc := false
	isDec := false
	curr := nums[0]
	i := 1
	for i < len(nums) {
		maxStreak = max(maxStreak, currStreak)

		if nums[i] > curr {
			if !isInc && !isDec {
				isInc = true
			}
			if isDec {
				isDec = false
				currStreak = 1
				continue
			}
			currStreak++
			curr = nums[i]
			i++
			continue
		}
		if nums[i] < curr {
			if !isInc && !isDec {
				isDec = true
			}
			if isInc {
				isInc = false
				currStreak = 1
				continue
			}
			currStreak++
			curr = nums[i]
			i++
			continue
		}
		isDec, isInc = false, false
		currStreak = 1
		curr = nums[i]
		i++
	}
	return max(maxStreak, currStreak)
}

func longestMonotonicSubarray(nums []int) int {
	incLen, decLen := 1, 1
	maxLen := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			incLen++
			decLen = 1
		} else if nums[i+1] < nums[i] {
			incLen = 1
			decLen++
		} else {
			incLen, decLen = 1, 1
		}
		maxLen = max(maxLen, incLen, decLen)
	}
	return maxLen
}

func main() {
	println(MyLongestMonotonicSubarray([]int{1, 4, 3, 3, 2})) // 2
	println(MyLongestMonotonicSubarray([]int{3, 3, 3, 3}))    // 1
	println(MyLongestMonotonicSubarray([]int{3, 2, 1}))       // 3
	println(longestMonotonicSubarray([]int{1, 4, 3, 3, 2}))   // 2
	println(longestMonotonicSubarray([]int{3, 3, 3, 3}))      // 1
	println(longestMonotonicSubarray([]int{3, 2, 1}))         // 3
}
