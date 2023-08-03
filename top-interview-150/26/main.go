package main

func removeDuplicates(nums []int) int {
	idx, idxToRemove := 0, 0
	for idxToRemove < len(nums) {
		if nums[idxToRemove] == nums[idx] {
			idxToRemove++
			continue
		}
		idx++
		tmp := nums[idx]
		nums[idx] = nums[idxToRemove]
		nums[idxToRemove] = tmp
		idxToRemove++
	}
	return idx + 1
}
func main() {
	nums := []int{1, 1, 2}
	println(removeDuplicates(nums))
}
