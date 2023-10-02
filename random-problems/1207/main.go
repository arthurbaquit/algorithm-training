package main

func uniqueOccurrences(arr []int) bool {
	occCount := make(map[int]int)
	for _, number := range arr {
		occCount[number]++
	}
	hasMultOcc := make(map[int]bool)
	for _, value := range occCount {
		_, ok := hasMultOcc[value]
		if ok {
			return false
		}
		hasMultOcc[value] = true
	}
	return true
}

func main() {
	println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	println(uniqueOccurrences([]int{1, 2}))
	println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
