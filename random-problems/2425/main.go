package main

func xorAllNums(nums1 []int, nums2 []int) int {
	ans := 0
	if len(nums1)%2 == 1 {
		for _, num2 := range nums2 {
			ans ^= num2
		}
	}
	if len(nums2)%2 == 1 {
		for _, num1 := range nums1 {
			ans ^= num1
		}
	}
	return ans
}

func main() {
	println(xorAllNums([]int{1, 2, 3}, []int{10, 2, 5, 0})) // 13
	println(xorAllNums([]int{1, 2}, []int{3, 4}))           // 0
}
