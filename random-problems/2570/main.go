package main

import "fmt"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	l1, l2 := 0, 0
	ans := [][]int{}
	for l1 < len(nums1) && l2 < len(nums2) {
		k1, v1 := nums1[l1][0], nums1[l1][1]
		k2, v2 := nums2[l2][0], nums2[l2][1]
		if k1 == k2 {
			ans = append(ans, []int{k1, v1 + v2})
			l1++
			l2++
			continue
		}
		if k1 > k2 {
			ans = append(ans, []int{k2, v2})
			l2++
			continue
		}
		ans = append(ans, []int{k1, v1})
		l1++
	}
	for i := l1; i < len(nums1); i++ {
		ans = append(ans, []int{nums1[i][0], nums1[i][1]})
	}
	for i := l2; i < len(nums2); i++ {
		ans = append(ans, []int{nums2[i][0], nums2[i][1]})
	}
	return ans
}

func main() {
	fmt.Println(mergeArrays([][]int{{1, 3}, {2, 2}, {3, 1}}, [][]int{{1, 2}, {2, 3}, {3, 4}})) // [[1,5],[2,5],[3,5]]
	fmt.Println(mergeArrays([][]int{{1, 1}}, [][]int{{1, 1}}))                                 // [[1,2]]
	fmt.Println(mergeArrays([][]int{{1, 1}}, [][]int{{2, 1}}))                                 // [[1,1],[2,1]]
}
