package main

func containsNearbyDuplicate(nums []int, k int) bool {
	visited := make(map[int][]int)
	for i, num := range nums {
		if _, ok := visited[num]; !ok {
			visited[num] = []int{i}
			continue
		}
		for _, v := range visited[num] {
			if absInt(v, i) <= k {
				return true
			}
		}
		visited[num] = append(visited[num], i)
	}
	return false
}

func absInt(x, y int) int {
	if x-y < 0 {
		return y - x
	}
	return x - y
}
func main() {
	println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
