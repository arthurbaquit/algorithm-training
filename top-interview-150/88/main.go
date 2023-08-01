package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	end, p1, p2 := m+n-1, m-1, n-1
	for e := end; p2 >= 0 && e > -1; e-- {
		if p1 > -1 && nums1[p1] > nums2[p2] {
			nums1[e] = nums1[p1]
			p1--
			continue
		}
		nums1[e] = nums2[p2]
		p2--
	}

}

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	merge(a, 3, b, 3)
	fmt.Println(a)
}
