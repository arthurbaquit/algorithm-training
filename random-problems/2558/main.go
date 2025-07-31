package main

import (
	"container/heap"
	"math"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
	h := &MaxHeap{}
	for _, gift := range gifts {
		h.Push(gift)
	}

	// Initialize the heap
	heap.Init(h)
	for range k {
		maxGift := heap.Pop(h).(int)
		heap.Push(h, int(math.Sqrt(float64(maxGift))))
	}
	var ans int64
	for _, val := range *h {
		ans += int64(val)
	}
	return ans
}

func main() {
	println(pickGifts([]int{25, 64, 9, 4, 100}, 4)) // 29
	println(pickGifts([]int{1, 1, 1, 1}, 4))        // 4
}
