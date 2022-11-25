package main

import "fmt"

func findLength(nums1 []int, nums2 []int) int {
	maxLen := -1
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			count := 0
			for aux := 0; i+aux < len(nums1) && j+aux < len(nums2) && nums1[i+aux] == nums2[j+aux]; aux++ {
				count++
			}
			if count > maxLen {
				maxLen = count
			}
		}
	}
	return maxLen

}

func main() {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	fmt.Println(findLength(nums1, nums2))
}
