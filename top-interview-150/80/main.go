package main

func removeDuplicates(nums []int) int {
	count, slow := 1, 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] == nums[fast-1] {
			count++
		} else {
			count = 1
		}
		if count <= 2 {
			nums[slow] = nums[fast]
			slow++

		}
	}
	return slow
}

func main() {
	println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}
