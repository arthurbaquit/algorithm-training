package main

import "sort"

type NumberContainers struct {
	idx         map[int]int
	numIdx      map[int][]int
	isNumSorted map[int]bool
}

func Constructor() NumberContainers {
	return NumberContainers{
		idx:         make(map[int]int),
		numIdx:      make(map[int][]int),
		isNumSorted: make(map[int]bool),
	}
}

func (nc *NumberContainers) Change(index int, number int) {
	v, ok := nc.idx[index]
	nc.idx[index] = number
	if ok {
		// rm idx from numIdx
		numIdxArr := nc.numIdx[v]
		for i, idx := range numIdxArr {
			if idx == index {
				nc.numIdx[v] = append(numIdxArr[:i], numIdxArr[i+1:]...)
				break
			}
		}
	}
	// add
	nc.numIdx[number] = append(nc.numIdx[number], index)
	nc.isNumSorted[number] = false
}

func (nc *NumberContainers) Find(number int) int {
	arr, ok := nc.numIdx[number]
	if !ok || len(arr) == 0 {
		return -1
	}
	if nc.isNumSorted[number] {
		return arr[0]
	}
	sort.Ints(nc.numIdx[number])
	nc.isNumSorted[number] = true
	return nc.numIdx[number][0]
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */

func main() {
	obj := Constructor()
	println(obj.Find(10)) // -1
	obj.Change(2, 10)
	obj.Change(1, 10)
	obj.Change(3, 10)
	obj.Change(5, 10)
	println(obj.Find(10)) // 1
	obj.Change(1, 20)
	println(obj.Find(10)) // 2
}
