package main

func findLHS(nums []int) int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	ans := 0
	for key, value := range countMap {
		if v, ok := countMap[key+1]; ok {
			ans = max(ans, value+v)
		}
	}
	return ans
}

func main() {
	println(findLHS([]int{1, 3, 2, 2, 5, 4, 3, 7})) // 5
	println(findLHS([]int{1, 2, 3, 4}))             // 2
	println(findLHS([]int{1, 1, 1, 1}))             // 0
}
