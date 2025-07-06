package main

type FindSumPairs struct {
	nums2Map map[int]int
	nums1    []int
	nums2    []int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	nums2Map := make(map[int]int)
	for _, num := range nums2 {
		nums2Map[num]++
	}
	return FindSumPairs{
		nums2Map: nums2Map,
		nums1:    nums1,
		nums2:    nums2,
	}
}

func (fsp *FindSumPairs) Add(index int, val int) {
	if index >= len(fsp.nums2) {
		return
	}
	fsp.nums2Map[fsp.nums2[index]]--
	fsp.nums2[index] += val
	fsp.nums2Map[fsp.nums2[index]]++
}

func (fsp *FindSumPairs) Count(tot int) int {
	count := 0
	for _, num := range fsp.nums1 {
		count += (fsp.nums2Map[tot-num])
	}
	return count
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */

func main() {
	obj := Constructor([]int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4})
	println(obj.Count(7)) // 8
	obj.Add(3, 2)
	println(obj.Count(8)) // 2
	println(obj.Count(4)) // 1
	obj.Add(0, 1)
	obj.Add(1, 1)
	println(obj.Count(7)) // 11
}
