package main

import (
	"container/heap"
	"sort"
)

func minOperations(nums []int, k int) int {
	sort.Ints(nums)
	count := 0
	for len(nums) > 1 {
		if nums[0] >= k {
			return count
		}
		count++
		numToAdd := nums[0]*2 + nums[1]
		idx := sort.SearchInts(nums, numToAdd)
		if idx == len(nums) {
			nums = append(nums[2:], numToAdd)
			continue
		}
		nums = append(nums[2:idx+1], nums[idx:]...)
		nums[idx-2] = numToAdd
	}
	return count
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func minOperationsHeap(a []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, val := range a {
		heap.Push(h, val)
	}
	count := 0
	for h.Len() > 1 {
		x := heap.Pop(h).(int)
		if x >= k {
			return count
		}
		count++
		y := heap.Pop(h).(int)
		t := int64(2*x + y) // test overflow
		if t < int64(k) {
			heap.Push(h, int(t))
		} else {
			heap.Push(h, k)
		}
	}
	return count
}

func main() {
	println(minOperations([]int{2, 11, 10, 1, 3}, 10))     // 2
	println(minOperationsHeap([]int{2, 11, 10, 1, 3}, 10)) // 2
	println(minOperations([]int{1, 1, 2, 4, 9}, 20))       // 4
	println(minOperationsHeap([]int{1, 1, 2, 4, 9}, 20))   // 4
}
