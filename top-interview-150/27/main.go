package main

func removeElement(nums []int, val int) int {
	idx, idxToRemove := 0, len(nums)-1
	for idx <= idxToRemove {
		if nums[idx] == val {
			tmp := nums[idx]
			nums[idx] = nums[idxToRemove]
			nums[idxToRemove] = tmp
			idxToRemove--
			continue
		}
		idx++
	}
	return idxToRemove + 1
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	println(removeElement(nums, val))
}
