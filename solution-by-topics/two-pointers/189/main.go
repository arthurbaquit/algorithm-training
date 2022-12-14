package main

import "fmt"

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	if k == 0 {
		return
	}
	count := 0
	for start := 0; count < l; start++ {
		curr, prev := start, nums[start]
		for {
			next_idx := (curr + k) % l
			temp := nums[next_idx]
			nums[next_idx] = prev
			prev = temp
			curr = next_idx
			count++
			if curr == start {
				break
			}
		}
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}
