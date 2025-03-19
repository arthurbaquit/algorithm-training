package main

func minOperations(nums []int) int {
	count := 0
	hasZero := false
	for i := range nums {
		if nums[i] == 0 {
			hasZero = false
			count++
			if i > len(nums)-3 {
				return -1
			}
			nums[i] = 1 - nums[i]
			nums[i+1] = 1 - nums[i+1]
			nums[i+2] = 1 - nums[i+2]
			if nums[i] == 0 || nums[i+1] == 0 || nums[i+2] == 0 {
				hasZero = true
			}
		}
	}
	if hasZero {
		return -1
	}
	return count
}

func main() {
	println(minOperations([]int{0, 1, 1, 1, 0, 0})) // 3
	println(minOperations([]int{0, 1, 1, 1}))       // -1
}
