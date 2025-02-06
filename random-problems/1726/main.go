package main

func tupleSameProduct(nums []int) int {
	prod := make(map[int]int)
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			prod[v*nums[j]]++
		}
	}
	ans := 0
	for _, v := range prod {
		ans += v * (v - 1) * 4
	}
	return ans
}

func main() {
	println(tupleSameProduct([]int{2, 3, 4, 6}))        // 8
	println(tupleSameProduct([]int{1, 2, 4, 5, 10}))    // 16
	println(tupleSameProduct([]int{2, 3, 4, 6, 8, 12})) // 40
}
