package main

func majorityElement(nums []int) int {
	freq := make(map[int]int)
	l := len(nums)
	for _, v := range nums {
		freq[v]++
		if freq[v] > l/2 {
			return v
		}
	}
	return 0
}

func main() {
	println(majorityElement([]int{3, 2, 3}))
	println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
