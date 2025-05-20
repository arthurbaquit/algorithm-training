package main

func isZeroArray(nums []int, queries [][]int) bool {
	accFreq := make([]int, len(nums))
	for _, query := range queries {
		l, r := query[0], query[1]
		accFreq[l]--
		if r == len(nums)-1 {
			continue
		}
		accFreq[r+1]++
	}
	acc := 0
	for i := range nums {
		acc += accFreq[i]
		nums[i] += acc
	}
	for _, num := range nums {
		if num > 0 {
			return false
		}
	}
	return true
}

func main() {
	println(isZeroArray([]int{1, 0, 1}, [][]int{{0, 2}}))            // true
	println(isZeroArray([]int{4, 3, 2, 1}, [][]int{{1, 3}, {0, 2}})) // false
}
