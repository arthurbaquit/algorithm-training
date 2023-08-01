package main

func removeDuplicates(nums []int) int {
	curr, idx := nums[0], 1
	for _, v := range nums {
		if v == curr {
			continue
		}
		nums[idx] = v
		idx++
		curr = v
	}
	return idx
}

func main() {
	println(removeDuplicates([]int{1, 1, 2}))
	println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
