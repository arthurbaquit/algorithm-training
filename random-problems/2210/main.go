package main

func countHillValley(nums []int) int {
	count := 0
	curr := 1
	for curr < len(nums)-1 {
		l, r := curr-1, curr+1
		for r < len(nums) && nums[r] == nums[curr] {
			r++
		}
		for l >= 0 && nums[l] == nums[curr] {
			l--
		}
		if l < 0 || r >= len(nums) {
			curr = r
			continue
		}
		isGTLeft := true
		if nums[curr] < nums[l] {
			isGTLeft = false
		}
		isGTRight := true
		if nums[curr] < nums[r] {
			isGTRight = false
		}
		if (isGTLeft && isGTRight) || (!isGTLeft && !isGTRight) {
			count += 1
		}
		curr = r
	}
	return count
}

func main() {
	println(countHillValley([]int{2, 4, 1, 1, 6, 5})) // 3
	println(countHillValley([]int{6, 6, 5, 5, 4, 1})) // 0
}
