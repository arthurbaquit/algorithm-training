package main

func isArraySpecial(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		parity := nums[i]%2 + nums[i+1]%2
		if parity != 1 {
			return false
		}
	}
	return true
}

func main() {
	println(isArraySpecial([]int{1, 2, 3}))    // true
	println(isArraySpecial([]int{2, 1, 4}))    // true
	println(isArraySpecial([]int{4, 3, 1, 6})) // false
}
