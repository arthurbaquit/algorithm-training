package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
    sort.Slice(arr, func(i, j int) bool {
        iCount, jCount := getBit1Count(arr[i]), getBit1Count(arr[j])
        if iCount == jCount {
            return arr[i] < arr[j]
        }
        return iCount < jCount
    })
    return arr
}

func getBit1Count(val int) (count int) {
    for val != 0 {
        count += val & 0x1
        val >>= 1
    }
    return
}

func main() {
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
	fmt.Println(sortByBits([]int{10000, 10000}))
	fmt.Println(sortByBits([]int{2, 3, 5, 7, 11, 13, 17, 19}))
	fmt.Println(sortByBits([]int{10, 100, 1000, 10000}))
}