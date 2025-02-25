package main

func numOfSubarrays(arr []int) int {
	modNum := 1000000000 + 7
	count, oddCount, evenCount, acc := 0, 0, 1, 0
	for _, v := range arr {
		acc += v % 2
		if acc%2 == 0 {
			count += oddCount
			evenCount++
			continue
		}
		count += evenCount
		oddCount++
	}
	return count % modNum
}

func main() {
	println(numOfSubarrays([]int{1, 3, 5}))             // 4
	println(numOfSubarrays([]int{2, 4, 6}))             // 0
	println(numOfSubarrays([]int{1, 2, 3, 4, 5, 6, 7})) // 16
	println(numOfSubarrays([]int{100, 100, 99, 99}))    // 4
}
