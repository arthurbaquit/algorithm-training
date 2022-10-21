package main

func containsNearbyDuplicate(nums []int, k int) bool {
	visited := make(map[int]int)
	for i, num := range nums {
		if _, ok := visited[num]; !ok {
			visited[num] = i
			continue
		}
		if i-visited[num] <= k {
			return true
		}
		visited[num] = i
	}
	return false
}

func main() {
	println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
