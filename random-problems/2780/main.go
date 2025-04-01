package main

func minimumIndex(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	currFreq := make(map[int]int)
	for i := range nums {
		num := nums[i]
		currFreq[num]++
		if currFreq[num] > (i+1)/2 && freq[num]-currFreq[num] > (len(nums)-i-1)/2 {
			return i
		}
	}
	return -1
}

func main() {
	println(minimumIndex([]int{1, 2, 2, 2}))                   // 2
	println(minimumIndex([]int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1})) // 4
	println(minimumIndex([]int{3, 3, 3, 3, 7, 2, 2}))          // -1
}
