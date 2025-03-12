package main

func maximumCount(nums []int) int {
	countN, countP, ans := 0, 0, 0
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if num > 0 {
			countP++
		} else {
			countN++
		}
		ans = max(countP, countN)
	}
	return ans
}

func main() {
	println(maximumCount([]int{-2, -1, -1, 1, 2, 3}))    // 3
	println(maximumCount([]int{-3, -2, -1, 0, 0, 1, 2})) // 3
	println(maximumCount([]int{5, 20, 66, 1314}))        // 4
}
