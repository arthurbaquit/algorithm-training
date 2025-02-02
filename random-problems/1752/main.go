package main

func check(nums []int) bool {
	needRotation := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			continue
		}
		if !needRotation {
			needRotation = true
			continue
		}
		return false
	}
	if !needRotation {
		return true
	}
	return nums[len(nums)-1] <= nums[0]
}

func main() {
	println(check([]int{3, 4, 5, 1, 2})) // true
	println(check([]int{2, 1, 3, 4}))    // false
	println(check([]int{1, 2, 3}))       // true
}
